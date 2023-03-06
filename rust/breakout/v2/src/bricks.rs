use crate::components::Brick;
use crate::{BrickSize, WinSize, BRICKS_COLUMNS, BRICKS_LINES};
use bevy::prelude::*;

pub struct BricksPlugin;

impl Plugin for BricksPlugin {
    fn build(&self, app: &mut App) {
        app.add_startup_system_to_stage(StartupStage::PostStartup, bricks_spawn_system);
    }
}

fn bricks_spawn_system(mut commands: Commands, win_size: Res<WinSize>) {
    let window_left = -win_size.w / 2.;
    let window_top = win_size.h / 2.;

    let wall_gap = 0.05 * win_size.w; // Gap between wall and first bricks
    let brick_gap = 0.01 * win_size.w; // Gap in between bricks
    let brick_size = compute_bricks_size(brick_gap, wall_gap, win_size.w, win_size.h);

    commands.insert_resource(BrickSize {
        w: brick_size.x,
        h: brick_size.y,
    });

    for l in 0..BRICKS_LINES {
        for c in 0..BRICKS_COLUMNS {
            let transform_x =
                window_left + brick_size.x / 2. + wall_gap + c as f32 * (brick_size.x + brick_gap);
            let transform_y =
                window_top - brick_size.y / 2. - wall_gap - l as f32 * (brick_size.y + brick_gap);

            commands
                .spawn(SpriteBundle {
                    sprite: Sprite {
                        color: Color::rgb(0.5, 0.5, 0.5),
                        custom_size: Some(brick_size),
                        ..default()
                    },
                    transform: Transform::from_xyz(transform_x, transform_y, 0.),
                    ..default()
                })
                .insert(Brick);
        }
    }
}

fn compute_bricks_size(brick_gap: f32, wall_gap: f32, win_w: f32, win_h: f32) -> Vec2 {
    let nb_columns = BRICKS_COLUMNS as f32;
    let nb_lines = BRICKS_LINES as f32;

    let bricks_grid_w = win_w - wall_gap * 2.; // Total width of grid
    let bricks_grid_h = win_h / 2. - wall_gap * 2.; // Total height of grid

    let grid_w_without_gaps = bricks_grid_w - brick_gap * (nb_columns - 1.);
    let grid_h_without_gaps = bricks_grid_h - brick_gap * (nb_lines - 1.);

    let brick_size_x = grid_w_without_gaps / nb_columns;
    let brick_size_y = grid_h_without_gaps / nb_lines;

    Vec2::new(brick_size_x, brick_size_y)
}
