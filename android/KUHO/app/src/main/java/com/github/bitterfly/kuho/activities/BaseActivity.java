package com.github.bitterfly.kuho.activities;

import android.content.Context;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.preference.PreferenceManager;
import android.support.v7.app.AppCompatActivity;

import com.github.bitterfly.kuho.R;

/**
 * Created by human on 8/14/16.
 */
public abstract class BaseActivity extends AppCompatActivity {
    final int layout;
    Context context;
    SharedPreferences preferences;

    public BaseActivity(int layout){
        this.layout = layout;
        this.context = context;
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(this.layout);

        loadUtils();
        bindWidgets();
        setListeners();
        loadData();
    }

    private void loadUtils() {
        preferences = PreferenceManager.getDefaultSharedPreferences(this);
    }

    protected abstract void bindWidgets();

    protected abstract void loadData();

    protected abstract void setListeners();

}
