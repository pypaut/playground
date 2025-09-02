use bevy::{color::palettes::basic::PURPLE, prelude::*, app::AppExit};

#[derive(Component)]
struct Player;

#[derive(Component)]
struct Ball;

#[derive(Component)]
pub struct Direction {
    pub x: f32,
    pub y: f32,
}

#[derive(Component)]
struct Buttons {
    up: KeyCode,
    down: KeyCode,
}

const BALL_RADIUS: f32 = 15.;
const BALL_SPEED: f32 = 300.;
const PLAYER_SIZE: f32 = 120.;
const PLAYER_SPEED: f32 = 500.;

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup)
        .add_systems(Update, players_handle_kb_events)
        .add_systems(Update, players_update)
        .add_systems(Update, players_handle_collision)
        .add_systems(Update, ball_handle_kb_events)
        .add_systems(Update, ball_update)
        .add_systems(Update, ball_handle_collision)
        .run();
}

fn setup(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<ColorMaterial>>,
) {
    commands.spawn(Camera2d);

    // Ball
    commands
        .spawn((
            Mesh2d(meshes.add(Circle::default())),
            MeshMaterial2d(materials.add(Color::from(PURPLE))),
            Transform::default().with_scale(Vec3::splat(BALL_RADIUS)),
        ))
        .insert(Ball)
        .insert(Direction { x: 0., y: 0. });

    // Player 1
    commands
        .spawn((
            Mesh2d(meshes.add(Rectangle::default())),
            MeshMaterial2d(materials.add(Color::from(PURPLE))),
            Transform::from_xyz(-500., 0., 0.)
                .with_scale(Vec3::new(15., PLAYER_SIZE, 0.)),
        ))
        .insert(Player)
        .insert(Direction { x: 0., y: 0. })
        .insert(Buttons {
            up: KeyCode::KeyW,
            down: KeyCode::KeyS,
        });

    // Player 2
    commands
        .spawn((
            Mesh2d(meshes.add(Rectangle::default())),
            MeshMaterial2d(materials.add(Color::from(PURPLE))),
            Transform::from_xyz(500., 0., 0.)
                .with_scale(Vec3::new(15., PLAYER_SIZE, 0.)),
        ))
        .insert(Player)
        .insert(Direction { x: 0., y: 0. })
        .insert(Buttons {
            up: KeyCode::KeyI,
            down: KeyCode::KeyK,
        });
}

fn players_update(
    time: Res<Time>,
    mut query: Query<(&Direction, &mut Transform), With<Player>>,
) {
    for (direction, mut transform) in query.iter_mut() {
        let translation = &mut transform.translation;
        translation.y += direction.y * time.delta_secs() * PLAYER_SPEED;
    }
}

fn players_handle_kb_events(
    kb: Res<ButtonInput<KeyCode>>,
    mut query: Query<(&mut Direction, &Buttons), With<Player>>,
) {
    for (mut direction, buttons) in query.iter_mut() {
        direction.y = if kb.pressed(buttons.up) {
            1.
        } else if kb.pressed(buttons.down) {
            -1.
        } else {
            0.
        }
    }
}

fn players_handle_collision(
    windows: Query<&mut Window>,
    mut query: Query<&mut Transform, With<Player>>,
) {
    let win_h = windows.single().unwrap().resolution.height();

    for mut transform in query.iter_mut() {
        let translation = &mut transform.translation;

        if translation.y - PLAYER_SIZE / 2. < -win_h / 2. {
            translation.y = -win_h / 2. + PLAYER_SIZE / 2.;
        } else if translation.y + PLAYER_SIZE / 2. > win_h / 2. {
            translation.y = win_h / 2. - PLAYER_SIZE / 2.;
        }
    }
}

fn ball_handle_collision(
    windows: Query<&mut Window>,
    mut exit: EventWriter<AppExit>,
    mut query: Query<(&mut Transform, &mut Direction), With<Ball>>,
) {
    let win_w = windows.single().unwrap().resolution.width();
    let win_h = windows.single().unwrap().resolution.height();

    for (mut transform, mut direction) in query.iter_mut() {
        let translation = &mut transform.translation;

        // Wall collision
        let top_collision = translation.y + BALL_RADIUS / 2. >= win_h / 2.;
        let bottom_collision = translation.y - BALL_RADIUS / 2. <= -win_h / 2.;
        if top_collision || bottom_collision {
            direction.y *= -1.;
        }

        let left_collision = translation.x + BALL_RADIUS / 2. > win_w / 2.;
        let right_collision = translation.x - BALL_RADIUS / 2. < -win_w / 2.;
        if left_collision || right_collision {
            exit.write(AppExit::Success);
        }
    }
}

fn ball_handle_kb_events(
    kb: Res<ButtonInput<KeyCode>>,
    mut query: Query<&mut Direction, With<Ball>>,
) {
    for mut direction in query.iter_mut() {
        if kb.just_pressed(KeyCode::Space)
            && direction.x == 0.
            && direction.y == 0.
        {
            // Start game
            direction.x = 0.5;
            direction.y = 0.5;
        }
    }
}

fn ball_update(
    time: Res<Time>,
    mut query: Query<(&Direction, &mut Transform), With<Ball>>,
) {
    for (direction, mut transform) in query.iter_mut() {
        let translation = &mut transform.translation;
        translation.x += direction.x * time.delta_secs() * BALL_SPEED;
        translation.y += direction.y * time.delta_secs() * BALL_SPEED;
    }
}
