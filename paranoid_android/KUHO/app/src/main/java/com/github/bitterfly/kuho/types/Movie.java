package com.github.bitterfly.kuho.types;

/**
 * Created by human on 8/14/16.
 */
public class Movie {
    private String title;
    private Integer year;

    public Movie(){
        title = "Star Wars";
        year = 1977;
    }

    public Movie(String title, Integer year){
        this.title = title;
        this.year = year;
    }

    public String getTitle() {
        return title;
    }

    public Integer getYear() {
        return year;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    @Override
    public String toString() {
        return "Movie{" +
                "title='" + title + '\'' +
                ", year=" + year +
                '}';
    }

    public void setYear(Integer year) {
        this.year = year;
    }
}
