#!/usr/bin/python3

import os
import sys
import http.server
import socketserver
import socket
import shutil
from base64 import b64encode
from urllib.parse import quote
from os.path import basename, splitext, join, isfile
from collections import defaultdict
from subprocess import run
from distutils.dir_util import copy_tree
from distutils.file_util import copy_file

build_dir = 'build'
source_dir = 'source'
dest_dir = 'built_static'

css_dir = join(build_dir, 'css')
images_dir = join(build_dir, 'images')

class TemporaryTCPServer(socketserver.TCPServer):
    def server_bind(self):
        self.socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        self.socket.bind(self.server_address)

def serve(port):
    os.chdir(dest_dir)
    handler = http.server.SimpleHTTPRequestHandler

    httpd = TemporaryTCPServer(("", port), handler)

    print("[serve] serving on port " + str(port))
    httpd.serve_forever()

def clean():
    shutil.rmtree(build_dir)
    shutil.rmtree(dest_dir)

def build():
    copy_tree(source_dir, build_dir, update=1)

    make_fallback_images(images_dir)

    print('[create] _images.scss ', end='')
    save_images_css(images_dir, join(css_dir, '_images.scss'))
    print('[ok]')

    run_sass(css_dir, join(dest_dir, 'css'))

    print('[update] asis ', end='')
    copy_tree(join(source_dir, 'asis'), join(dest_dir, 'asis'), update=1)
    print('[ok]')

def run_sass(css_source_dir, css_dest_dir):
    os.makedirs(css_dest_dir, exist_ok=True)

    for (dirpath, dirnames, filenames) in os.walk(css_source_dir):
        for f in filenames:
            name, ext = splitext(f)
            if ext == '.scss' and name[0] != '_':
                print("[sass] " + f + ' ', end='')
                run([
                    'sass',
                    join(css_source_dir, f),
                    join(css_dest_dir, name + '.css')
                    ], check = True)
                print("[ok]")
            elif ext == '.css':
                print("[copy] " + f + ' ', end='')
                copy_file(join(css_source_dir, f), join(css_dest_dir, f), update=1)
                print("[ok]")
        break

def make_fallback_images(images_dir):
    images = find_built_images(images_dir)
    for image, files in images.items():
        f = files[0]

        pngimage = image + '.png'
        if pngimage not in files:
            print("[create] " + pngimage + ' ', end='')
            run([
                'convert', 
                '-background', 'none',
                join(images_dir, f), 
                join(images_dir, pngimage)
            ], check = True)
            print("[ok]")


def images_in_dir(dir):
    vectors = []
    rasters = []
    dumb_rasters = []
    lossy = []
    for (dirpath, dirnames, filenames) in os.walk(dir):
        for f in filenames:
            name, ext = splitext(basename(f))
            if ext in ['.svg']:
                vectors += [f]
            if ext in ['.png']:
                rasters += [f]
            if ext in ['.gif']:
                dumb_rasters += [f]
            if ext in ['.jpg', '.jpeg']:
                lossy += [f]

        break

    return vectors + rasters + dumb_rasters + lossy

def find_built_images(images_dir):
    images = defaultdict(list)

    for image in images_in_dir(images_dir):
        name, _ = splitext(basename(image))
        images[name] += [image]

    return dict(images)

def images_to_css(images_dir):
    images = find_built_images(images_dir)
    csseses = []

    for name, files in images.items():
        css = '.image-' + name + " {\n"
        files_and_extensions = [(f, splitext(f)[1][1:]) for f in files]

        for image, ext in [(f, ext) for f, ext in files_and_extensions if ext != 'svg']:
            data = raster_data(join(images_dir, image), ext)
            css += 'background-image: url(' + data + ");\n"

        for svg, ext in [(f, ext) for f, ext in files_and_extensions if ext == 'svg']:
            data = xml_data(join(images_dir, svg), ext)
            css += 'background-image: url(' + data + "), linear-gradient(transparent, transparent);\n"

        css += "}\n"
        csseses += [css]

    return "\n".join(csseses)

def save_images_css(images_dir, css_file):
    with open(css_file, 'w') as f:
        f.write(images_to_css(images_dir))

def raster_data(image_filename, ext):
    with open(image_filename, 'rb') as f:
        data = b64encode(f.read()).decode('utf-8')
        return 'data:image/' + ext + ';base64,' + data

def xml_data(image_filename, ext):
    with open(image_filename, 'r') as f:
        data = quote(f.read())
        return 'data:image/' + ext + '+xml;charset=US-ASCII,' + data

def image_data(image_filename):
    _, ext = splitext(image_filename)
    if ext == '.svg':
        return xml_data(image_filename, ext)
    else:
        return raster_data(image_filename, ext)

if __name__ == '__main__':
    try:
        arg = sys.argv[1]
    except IndexError:
        arg = None
    
    if arg == 'build':
        build()
    elif arg == 'clean':
        clean()
    elif arg == 'serve':
        try:
            port = int(sys.argv[2])
        except IndexError:
            port = 8000
        build()
        serve(port)
    else:
        print('please use "build", "clean" or "serve" as a first argument.')
