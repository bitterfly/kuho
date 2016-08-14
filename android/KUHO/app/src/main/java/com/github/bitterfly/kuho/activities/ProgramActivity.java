package com.github.bitterfly.kuho.activities;

import android.content.Context;
import android.content.Intent;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.ListView;

import com.github.bitterfly.kuho.R;
import com.github.bitterfly.kuho.adapters.MovieAdapter;
import com.github.bitterfly.kuho.types.Movie;

import java.util.Arrays;
import java.util.List;


public class ProgramActivity extends BaseActivity {
    ListView listView;
    Button button;
    private Button test;

    public ProgramActivity() {
        super(R.layout.activity_program);
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
    }

    @Override
    protected void loadData() {
        List<Movie> items = Arrays.asList(new Movie(), new Movie("My socks", 2016));

        ArrayAdapter data = new MovieAdapter(ProgramActivity.this, R.layout.adapter_movies, items);
        listView.setAdapter(data);
    }

    protected void bindWidgets() {
        this.button = (Button)findViewById(R.id.buttonShowSettings);
        this.test = (Button)findViewById(R.id.buttonTest);
        this.listView = (ListView)findViewById(R.id.listView);
    }

    protected void setListeners() {
        button.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                Intent intent = new Intent(ProgramActivity.this, SettingsActivity.class);
                startActivity(intent);
            }
        });

        test.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                Log.i("info: ", preferences.getString("server_address", ""));
            }
        });
    }

}
