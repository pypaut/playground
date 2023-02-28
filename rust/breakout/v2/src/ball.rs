use crate::components::{Ball, Direction, Player};
use crate::{WinSize, BALL_BASE_SPEED, BALL_SIZE, PLAYER_SIZE, TIME_STEP};
use bevy::{prelude::*, sprite::MaterialMesh2dBundle};

pub struct BallPlugin;

impl Plugin for BallPlugin {
    fn build(&self, app: &mut App) {
        app.add_startup_system_to_stage(StartupStage::PostStartup, ball_spawn_system);
        app.add_system(ball_movement_system);
        app.add_system(ball_wall_collision_system);
        app.add_system(ball_player_collision_system);
    }
}

fn ball_spawn_system(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<ColorMaterial>>,
) {
    commands
        .spawn(MaterialMesh2dBundle {
            mesh: meshes.add(shape::Circle::new(BALL_SIZE).into()).into(),
            material: materials.add(ColorMaterial::from(Color::PURPLE)),
            transform: Transform::from_xyz(0., 0., 0.),
            ..default()
        })
        .insert(Ball)
        .insert(Direction { x: -0.5, y: -1. });
}

fn ball_movement_system(mut query: Query<(&mut Direction, &mut Transform), With<Ball>>) {
    for (direction, mut transform) in query.iter_mut() {
        let translation = &mut transform.translation;
        translation.x += direction.x * TIME_STEP * BALL_BASE_SPEED;
        translation.y += direction.y * TIME_STEP * BALL_BASE_SPEED;
    }
}

fn ball_wall_collision_system(
    win_size: Res<WinSize>,
    mut query: Query<(&mut Direction, &Transform), With<Ball>>,
) {
    for (mut direction, transform) in query.iter_mut() {
        let next_x = transform.translation.x + direction.x * TIME_STEP * BALL_BASE_SPEED;
        let next_y = transform.translation.y + direction.y * TIME_STEP * BALL_BASE_SPEED;

        // Ball reaches bottom or top
        if next_y - BALL_SIZE / 2. < -win_size.h / 2. || next_y + BALL_SIZE / 2. > win_size.h / 2. {
            direction.y *= -1.;
        }

        // Ball reaches left or right
        if next_x - BALL_SIZE / 2. < -win_size.w / 2. || next_x + BALL_SIZE / 2. > win_size.w / 2. {
            direction.x *= -1.;
        }

        direction.normalize();
    }
}

fn ball_player_collision_system(
    mut ball_query: Query<(&mut Direction, &Transform), With<Ball>>,
    mut player_query: Query<&Transform, With<Player>>,
) {
    for (mut ball_dir, ball_transform) in ball_query.iter_mut() {
        let next_ball_x = ball_transform.translation.x + ball_dir.x * TIME_STEP * BALL_BASE_SPEED;
        let next_ball_y = ball_transform.translation.y + ball_dir.y * TIME_STEP * BALL_BASE_SPEED;

        if let Ok(player_transform) = player_query.get_single_mut() {
            let player_x = player_transform.translation.x;
            let player_y = player_transform.translation.y;

            let player_left = player_x - PLAYER_SIZE.0 / 2.;
            let player_right = player_x + PLAYER_SIZE.0 / 2.;

            let player_bot = player_y - PLAYER_SIZE.1 / 2.;
            let player_top = player_y + PLAYER_SIZE.1 / 2.;

            // Check ball/player collision
            let x_cond = player_left < next_ball_x && next_ball_x < player_right;
            let y_cond = player_bot < next_ball_y && next_ball_y < player_top;
            if x_cond && y_cond {
                let c = (ball_transform.translation.x - player_x) / PLAYER_SIZE.0;

                ball_dir.y *= -1.;
                ball_dir.x = c;

                ball_dir.normalize();
            }
        }
    }
}
