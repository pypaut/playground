use crate::ball::BallPlugin;
use crate::player::PlayerPlugin;
use bevy::prelude::*;

mod ball;
mod components;
mod player;

// region:    --- Constants

const PLAYER_SIZE: (f32, f32) = (200., 20.);
const PLAYER_BASE_SPEED: f32 = 800.;

const BALL_SIZE: f32 = 10.;
const BALL_BASE_SPEED: f32 = 500.;

// endregion: --- Constants

// region:    --- Resources

#[derive(Resource)]
pub struct WinSize {
    pub w: f32,
    pub h: f32,
}

// endregion: --- Resources

fn main() {
    App::new()
        .insert_resource(ClearColor(Color::rgb(0.04, 0.04, 0.04)))
        .add_plugins(DefaultPlugins.set(WindowPlugin {
            window: WindowDescriptor {
                title: "Breakout".to_string(),
                width: 1000.,
                height: 800.,
                // mode: WindowMode::BorderlessFullscreen,
                ..default()
            },
            ..default()
        }))
        .add_plugin(PlayerPlugin)
        .add_plugin(BallPlugin)
        .add_startup_system(setup_system)
        .run()
}

fn setup_system(mut commands: Commands, mut windows: ResMut<Windows>) {
    // Camera
    commands.spawn(Camera2dBundle::default());

    // Capture window size
    let window = windows.get_primary_mut().unwrap();
    let win_size = WinSize {
        w: window.width(),
        h: window.height(),
    };
    commands.insert_resource(win_size);
}
