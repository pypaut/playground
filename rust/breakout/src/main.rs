use bevy::{prelude::*, sprite::MaterialMesh2dBundle};

use crate::ball::Ball;
use crate::bricks::Bricks;
use crate::player::Player;

mod ball;
mod bricks;
mod player;

fn main() {
    App::new()
        .insert_resource(ClearColor(Color::rgb(0.04, 0.04, 0.04)))
        .add_plugins(DefaultPlugins.set(WindowPlugin {
            window: WindowDescriptor {
                width: 1920.,
                height: 1080.,
                title: "Breakout".to_string(),
                ..default()
            },
            ..default()
        }))
        .add_startup_system(setup_system)
        .run();
}

fn setup_system(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<ColorMaterial>>,
) {
    commands.spawn(Camera2dBundle::default());

    // Bar (player)
    let player = Player::new(
        Color::rgb(0.25, 0.25, 0.75),
        Vec2::new(200.0, 20.0),
        Transform::from_xyz(0., -400., 0.),
    );

    player.spawn(&mut commands);

    // Bricks
    let brick_size = Vec2::new(100.0, 50.0);
    let start_x = -880.0;
    let start_y = 500.0;
    let brick_gap = 10.0;
    let nb_lines = 5;
    let nb_columns = 17;

    let bricks = Bricks::new(
        Color::rgb(1., 1., 1.),
        brick_size,
        start_x,
        start_y,
        brick_gap,
        nb_lines,
        nb_columns,
    );

    bricks.spawn(&mut commands);

    // Ball
    let ball = Ball::new(
        &mut commands,
        &mut meshes,
        &mut materials,
        10.,
        Color::PURPLE,
        Transform::from_translation(Vec3::new(0., 0., 0.)),
    );
}
