use crate::components::{Direction, Player};
use crate::{WinSize, PLAYER_BASE_SPEED, PLAYER_SIZE};
use bevy::prelude::*;

pub struct PlayerPlugin;

impl Plugin for PlayerPlugin {
    fn build(&self, app: &mut App) {
        app.add_startup_system_to_stage(StartupStage::PostStartup, player_spawn_system);
        app.add_system(player_keyboard_event_system);
        app.add_system(player_movement_system.after(player_keyboard_event_system));
    }
}

fn player_spawn_system(mut commands: Commands, win_size: Res<WinSize>) {
    let bottom = -win_size.h / 2.;
    commands
        .spawn(SpriteBundle {
            sprite: Sprite {
                color: Color::rgb(0.25, 0.25, 0.75),
                custom_size: Some(Vec2::new(PLAYER_SIZE.0, PLAYER_SIZE.1)),
                ..default()
            },
            transform: Transform::from_xyz(0., bottom + PLAYER_SIZE.1 / 2. + 50., 0.),
            ..default()
        })
        .insert(Player)
        .insert(Direction { x: 0., y: 0. });
}

fn player_keyboard_event_system(
    kb: Res<Input<KeyCode>>,
    mut query: Query<&mut Direction, With<Player>>,
) {
    if let Ok(mut direction) = query.get_single_mut() {
        direction.x = if kb.pressed(KeyCode::A) {
            -1.
        } else if kb.pressed(KeyCode::D) {
            1.
        } else {
            0.
        }
    }
}

fn player_movement_system(
    time: Res<Time>,
    win_size: Res<WinSize>,
    mut query: Query<(&Direction, &mut Transform), With<Player>>,
) {
    for (direction, mut transform) in query.iter_mut() {
        let translation = &mut transform.translation;

        let player_left = translation.x - PLAYER_SIZE.0 / 2.;
        let player_right = translation.x + PLAYER_SIZE.0 / 2.;

        let screen_left = -win_size.w / 2.;
        let screen_right = win_size.w / 2.;

        let is_in_bound = screen_left < player_left && player_right < screen_right;
        let is_left_toward_right = player_left < screen_left && direction.x > 0.;
        let is_right_toward_left = screen_right < player_right && direction.x < 0.;

        if is_in_bound || is_left_toward_right || is_right_toward_left {
            translation.x += direction.x * time.delta_seconds() * PLAYER_BASE_SPEED;
        }
    }
}
