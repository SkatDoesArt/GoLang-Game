

# GoLang-Game
A 2D Go game based on Ebitengine (Ebiten) illustrating terrain rendering via quadtree, map loading from files, random map generation, multiple camera modes, and various gameplay extensions (portals, zoom, save, treasure hunt, etc.).


<img src=goGamePic.png alt="GoGamePic">

---

## Quick Overview
- **Engine**: Ebitengine v2
- **Floor Rendering**: Simple grid, from file, via quadtree, or random generation (quadtree)
- **Movement**: Keyboard arrows, configurable animation
- **Camera**: Static, follow character, or limited to borders
- **Options**: Torus world (round earth), water blocking, animated tiles, zoom, portals, random map save, treasure hunt
- **Debug Overlay**: Grid, camera/character coordinates, etc. (press D)

---

## Prerequisites
- Go 1.21+ (tested with Go 1.23.2)
- Windows, macOS, or Linux. Commands below are for Windows PowerShell.

---

## Install Dependencies
From the repository root:
```powershell
go mod tidy
```

---

## Run the Game
Several options depending on your current directory:

- From the project root (targeting the app in `cmd/`):
```powershell
go run ./cmd -config ./cmd/config.json
```
- Or compile and run:
```powershell
go build ./cmd
./cmd.exe -config ./cmd/config.json
```
**Notes:**
- The `-config` parameter is optional. Without it, the executable looks for `config.json` in the current directory.
- Relative paths (e.g., to a map) are resolved from your current directory at runtime.

---

## Keyboard Controls
- **Arrow keys**: Move character
- **D**: Toggle debug display
- **Numpad +**: Zoom out (show more tiles) if `Zoom` is true
- **Numpad -**: Zoom in (show fewer tiles) if `Zoom` is true
- **T**: Place a portal (if `Portal` is true). Max two portals; stepping on one teleports to the other.
- **S**: Save the randomly generated map (if `FloorKind` = 3 and `Sauvegarde` is true)

---

## Configuration
The `cmd/config.json` file contains all settings. Example provided:
```json
{
	"DebugMode": false,
	"NumTileX": 9,
	"NumTileY": 9,
	"TileSize": 16,
	"NumCharacterAnimImages": 5,
	"NumFramePerCharacterAnimImage": 5,
	"NumTileForDebug": 6,
	"CameraMode": 1,
	"FloorKind": 2,
	"FloorFile": "../floor-files/beaupasbeau",
	"RandomMapXSize": 15,
	"RandomMapYSize": 15,
	"TerreRonde": false,
	"PasMarcheEau": false,
	"BlocAnimated": false,
	"Zoom": false,
	"Portal": false,
	"Sauvegarde": false,
	"Chasseautresor": false
}
```
**Main fields:**
- `DebugMode` (bool): Enable debug overlay (can also be toggled via D)
- `NumTileX`, `NumTileY` (int): Number of tiles visible on screen (excluding debug area)
- `TileSize` (int): Size (px) of a tile
- `NumCharacterAnimImages` (int): Number of images per character animation step
- `NumFramePerCharacterAnimImage` (int): Number of updates between animation frames
- `NumTileForDebug` (int): Width in tiles of the debug info area (right/top)
- `CameraMode` (int): 0=static, 1=follow character, 2=limited to borders
  - In mode 2, if `TerreRonde` is true, the camera follows the character (no border blocking)
- `FloorKind` (int): 0=grid, 1=from file, 2=quadtree (from file), 3=random generation (quadtree)
- `FloorFile` (string): Path to a map (used for `FloorKind` 1 or 2)
- `RandomMapXSize`, `RandomMapYSize` (int): Dimensions of the random map (for `FloorKind` 3)
- `TerreRonde` (bool): Torus world (horizontal/vertical wrap) for quadtree/file reading
- `PasMarcheEau` (bool): Prevent walking on water (tile type 4)
- `BlocAnimated` (bool): Enable animation for certain floor tiles
- `Zoom` (bool): Allow numpad +/− for zooming
- `Portal` (bool): Allow portal placement and teleportation (press T)
- `Sauvegarde` (bool): Allow saving a random map (press S)
- `Chasseautresor` (bool): Enable treasure hunt mode

**Notes:**
- Tile types: Floor tile indices come from `assets/floor.png`. Value 4 corresponds to water (used by `PasMarcheEau`).
- In "treasure hunt" mode, a random treasure is placed; the distance in blocks is displayed. At distance ≤ 1, the message "Treasure found! Victory!" appears.

---

## Project Structure
```
assets/         # Embedded images (go:embed) and credits
camera/         # Camera modes (static, follow, limited to borders)
character/      # Character animation and movement logic
cmd/            # Entry point (main), default config
configuration/  # Global configuration reading/storage
floor/          # Floor management (grid, file, quadtree, random) + drawing + collisions
floor-files/    # Example maps (text files with numbers)
game/           # Ebiten game loop (Init, Update, Draw, Layout)
portal/         # Portals: creation, update, drawing, teleportation
quadtree/       # Quadtree structure + construction/queries + tests
tresor/         # Treasure hunt mode (placement, drawing, distance)
```

---

## Tests
Tests cover quadtree construction and queries, as well as map reading.
Run all tests:
```powershell
go test ./...
```
---

## 🤝 Contributing

This is a personal project. If you’d like to suggest improvements:
- Open an issue describing the idea or bug
- If submitting a PR, keep changes focused and include a brief summary

---

## Troubleshooting
- **Black window/partial screen**: Ensure `assets.Load()` is called (it is in `cmd/main.go`) and that `TileSize`, `NumTileX/Y` result in a displayable window.
- **Map not found**: Adjust `FloorFile` according to your current directory. From the root, use a path like `./floor-files/map-name`.
- **Character "stuck" at the edge**: Switch `CameraMode` to 2 to limit the camera to borders, or enable `TerreRonde` for wrapping.
- **Zoom inactive**: Set `Zoom` to true and use the numpad (+ to zoom out, − to zoom in).
- **Portals inactive**: Set `Portal` to true. Max two portals; the first is replaced if a third is placed.
- **Save inactive**: Requires `FloorKind` = 3 (random map) and `Sauvegarde` = true. The file is written to `floor-files/` with a timestamp.

---
