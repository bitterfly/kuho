package com.github.bitterfly.kuho.activities;

import android.os.Bundle;
import android.preference.PreferenceFragment;

import com.github.bitterfly.kuho.R;

public class SettingsFragment extends PreferenceFragment {
    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        // Load the preferences from an XML resource
        addPreferencesFromResource(R.xml.preferences);
    }
}