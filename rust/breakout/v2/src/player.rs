use crate::components::{Direction, Player};
use crate::{WinSize, PLAYER_BASE_SPEED, PLAYER_SIZE, TIME_STEP};
use bevy::prelude::*;

pub struct PlayerPlugin;

impl Plugin for PlayerPlugin {
    fn build(&self, app: &mut App) {
        app.add_startup_system_to_stage(StartupStage::PostStartup, player_spawn_system);
        app.add_system(player_keyboard_event_system);
        app.add_system(player_movement_system);
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

fn player_movement_system(mut query: Query<(&Direction, &mut Transform), With<Player>>) {
    for (direction, mut transform) in query.iter_mut() {
        let translation = &mut transform.translation;
        translation.x += direction.x * TIME_STEP * PLAYER_BASE_SPEED;
    }
}
