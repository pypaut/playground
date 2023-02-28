use crate::components::{Ball, Direction};
use crate::{WinSize, BALL_BASE_SPEED, BALL_SIZE, TIME_STEP};
use bevy::{prelude::*, sprite::MaterialMesh2dBundle};

pub struct BallPlugin;

impl Plugin for BallPlugin {
    fn build(&self, app: &mut App) {
        app.add_startup_system_to_stage(StartupStage::PostStartup, ball_spawn_system);
        app.add_system(ball_movement_system);
        app.add_system(ball_collision_system);
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
        .insert(Direction { x: 0., y: -1. });
}

fn ball_movement_system(mut query: Query<(&mut Direction, &mut Transform), With<Ball>>) {
    for (mut direction, mut transform) in query.iter_mut() {
        let translation = &mut transform.translation;
        translation.x += direction.x * TIME_STEP * BALL_BASE_SPEED;
        translation.y += direction.y * TIME_STEP * BALL_BASE_SPEED;
    }
}

fn ball_collision_system(
    win_size: Res<WinSize>,
    mut query: Query<(&mut Direction, &Transform), With<Ball>>,
) {
    for (mut direction, mut transform) in query.iter_mut() {
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
