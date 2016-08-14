package com.github.bitterfly.kuho.adapters;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import com.github.bitterfly.kuho.R;
import com.github.bitterfly.kuho.types.Movie;

import java.util.List;

/**
 * Created by human on 8/14/16.
 */
public class MovieAdapter extends ArrayAdapter<Movie> {

    private final int resource;

    public MovieAdapter(Context context, int resource, List objects) {
        super(context, resource, objects);
        this.resource = resource;
    }
    @Override
    public View getView(int position, View convertView, ViewGroup parent) {

        View v = convertView;

        if (v == null) {
            LayoutInflater vi;
            vi = LayoutInflater.from(getContext());
            v = vi.inflate(resource, null);
        }

        Movie movie = getItem(position);

        if (movie != null) {
            TextView title = (TextView) v.findViewById(R.id.adapterTitle);
            TextView year = (TextView) v.findViewById(R.id.adapterYear);

            if (title != null) {
                title.setText(String.format("\"%s\"", movie.getTitle()));
            }


            if (year != null) {
                year.setText(String.format("- %d", movie.getYear()));
            }
        }

        return v;
    }


}
