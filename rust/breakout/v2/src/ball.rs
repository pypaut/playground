use crate::components::{Ball, Brick, Direction, Player};
use crate::{BrickSize, WinSize, BALL_BASE_SPEED, BALL_SIZE, PLAYER_SIZE};
use bevy::{
    prelude::*,
    sprite::collide_aabb::{collide, Collision},
    sprite::MaterialMesh2dBundle,
};

pub struct BallPlugin;

impl Plugin for BallPlugin {
    fn build(&self, app: &mut App) {
        app.add_startup_system_to_stage(StartupStage::PostStartup, ball_spawn_system);
        app.add_system(ball_player_collision_system.after(ball_movement_system));
        app.add_system(ball_wall_collision_system.after(ball_player_collision_system));
        app.add_system(ball_brick_collision_system);
        app.add_system(ball_movement_system);
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

fn ball_movement_system(
    time: Res<Time>,
    mut query: Query<(&mut Direction, &mut Transform), With<Ball>>,
) {
    for (direction, mut transform) in query.iter_mut() {
        let translation = &mut transform.translation;
        translation.x += direction.x * time.delta_seconds() * BALL_BASE_SPEED;
        translation.y += direction.y * time.delta_seconds() * BALL_BASE_SPEED;
    }
}

fn ball_wall_collision_system(
    time: Res<Time>,
    win_size: Res<WinSize>,
    mut query: Query<(&mut Direction, &Transform), With<Ball>>,
) {
    for (mut direction, transform) in query.iter_mut() {
        let next_x = transform.translation.x + direction.x * time.delta_seconds() * BALL_BASE_SPEED;
        let next_y = transform.translation.y + direction.y * time.delta_seconds() * BALL_BASE_SPEED;

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
    time: Res<Time>,
    mut ball_query: Query<(&mut Direction, &Transform), With<Ball>>,
    mut player_query: Query<&Transform, With<Player>>,
) {
    for (mut ball_dir, ball_transform) in ball_query.iter_mut() {
        let next_ball_x =
            ball_transform.translation.x + ball_dir.x * time.delta_seconds() * BALL_BASE_SPEED;
        let next_ball_y =
            ball_transform.translation.y + ball_dir.y * time.delta_seconds() * BALL_BASE_SPEED;

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
                let c = 1.5 * (ball_transform.translation.x - player_x) / PLAYER_SIZE.0;

                ball_dir.y *= -1.;
                ball_dir.x = c;

                ball_dir.normalize();
            }
        }
    }
}

fn ball_brick_collision_system(
    mut commands: Commands,
    brick_size: Res<BrickSize>,
    time: Res<Time>,
    mut ball_query: Query<(&mut Direction, &Transform), With<Ball>>,
    mut bricks_query: Query<(Entity, &Transform), With<Brick>>,
) {
    for (mut ball_dir, ball_transform) in ball_query.iter_mut() {
        let next_ball_x =
            ball_transform.translation.x + ball_dir.x * time.delta_seconds() * BALL_BASE_SPEED;
        let next_ball_y =
            ball_transform.translation.y + ball_dir.y * time.delta_seconds() * BALL_BASE_SPEED;

        for (entity, brick_transform) in bricks_query.iter_mut() {
            let collision = collide(
                ball_transform.translation,
                ball_transform.scale.truncate(),
                brick_transform.translation,
                Vec2::new(brick_size.w, brick_size.h),
            );

            if collision.is_some() {
                let brick_x = brick_transform.translation.x;
                let brick_y = brick_transform.translation.y;

                let brick_left = brick_x - brick_size.w / 2.;
                let brick_right = brick_x + brick_size.w / 2.;

                let brick_bot = brick_y - brick_size.h / 2.;
                let brick_top = brick_y + brick_size.h / 2.;

                let ball_left = next_ball_x - BALL_SIZE / 2.;
                let ball_right = next_ball_x + BALL_SIZE / 2.;

                let ball_bot = next_ball_y - BALL_SIZE / 2.;
                let ball_top = next_ball_y + BALL_SIZE / 2.;

                // Vertical collision
                if ball_top > brick_bot || ball_bot < brick_top {
                    ball_dir.y *= -1.;
                }

                // Horizontal collision
                if ball_right > brick_left || ball_left < brick_right {
                    ball_dir.x *= -1.;
                }

                commands.entity(entity).despawn();
            }
        }
    }
}
