import { fromFile } from 'gen-readlines';

let sizeX;
let sizeY = 0;
let guard;
const obstacles = {};
let possibleObstacles = 0;

for (const line of fromFile(`input.txt`)) {
  const l = line.toString();
  if (!sizeX) sizeX = l.length;

  for (let j = 0; j < l.length; j++) {
    const vec = { x: j, y: sizeY };

    if (l[j] === '^') {
      guard = { ...vec, dir: '^', path: [] };
    }

    if (l[j] === '#') {
      if (obstacles[vec.x]) obstacles[vec.x].push(vec.y);
      else obstacles[vec.x] = [vec.y];
    }
  }

  sizeY++;
}

function isOutOfBounds({ x, y }) {
  return x < 0 || x >= sizeX || y < 0 || y >= sizeY;
}

function containsObstacle({ x, y }, o) {
  return o[x] && o[x].includes(y);
}

function getNext(g) {
  if (g.dir === '^') return { ...g, y: g.y - 1 };
  if (g.dir === '>') return { ...g, x: g.x + 1 };
  if (g.dir === 'v') return { ...g, y: g.y + 1 };
  if (g.dir === '<') return { ...g, x: g.x - 1 };
}

function detectLoop(g) {
  return g.path.length > sizeX * sizeY;
}

function turn(g) {
  if (g.dir === '^') g.dir = '>';
  else if (g.dir === '>') g.dir = 'v';
  else if (g.dir === 'v') g.dir = '<';
  else if (g.dir === '<') g.dir = '^';
}

function recordPath(g) {
  g.path.push({ x: guard.x, y: guard.y });
}

function move(g, o) {
  recordPath(g);

  const next = getNext(g);

  if (isOutOfBounds(next)) {
    return false;
  }

  if (detectLoop(g)) {
    throw new Error('Loop detected!');
  }

  if (containsObstacle(next, o)) {
    turn(g);
  } else {
    g.x = next.x;
    g.y = next.y;
  }

  return true;
}

function placeObstacle({ x, y }, o) {
  if (x === guard.x && y === guard.y) return;
  if (containsObstacle({ x, y }, o)) return;

  if (o[x]) o[x].push(y);
  else o[x] = [y];

  return true;
}

for (let x = 0; x < sizeX; x++) {
  for (let y = 0; y < sizeY; y++) {
    const g = JSON.parse(JSON.stringify(guard));
    const o = JSON.parse(JSON.stringify(obstacles));

    placeObstacle({ x, y }, o);

    try {
      while (move(g, o)) {
        /**/
      }
    } catch {
      possibleObstacles++;
    }
  }
}

console.log('Obstacle position possibilities:', possibleObstacles);