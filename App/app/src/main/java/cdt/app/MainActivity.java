package cdt.app;

import android.content.Intent;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;

public class MainActivity extends AppCompatActivity {


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        // Button to join a party
        final Button joinButton = findViewById(R.id.join_button_id);
        joinButton.setOnClickListener(new View.OnClickListener() {
            public void onClick(View v) {
                startActivity(new Intent(MainActivity.this, JoinActivity.class));
            }
        });

        // button to host a party
        final Button hostButton = findViewById(R.id.host_button_id);
        hostButton.setOnClickListener(new View.OnClickListener() {
            public void onClick(View v) {
                startActivity(new Intent(MainActivity.this, HostActivity.class));

            }
        });

        // button to settings activity
        final Button settingsButton = findViewById(R.id.settingsButton_id);
        settingsButton.setOnClickListener(new View.OnClickListener() {
            public void onClick(View v) {
                // Code here executes on main thread after user presses button

                   startActivity(new Intent(MainActivity.this, SettingsActivity.class));

            }
        });

    }
}




