import {
  redrawPlot,
  redrawGradientPlot,
  addSurfaceToPlot,
  addPointToGradientPlot
} from './plots.js';
import { updatePedictionFunction } from './parseFormat.js';

const spanElemento = document.getElementById('costValue');

export async function start(input, {
  maxIterations = 100,   // how many steps to take
  delay = 200           // ms between each update
} = {}) {
  const featuresMatrix = featuresMatrixToJs(input);
  // initial random data & surface
  const output = generateRandomDots(featuresMatrix);
  redrawPlot(output.x, output.y, output.z);
  let costSurface = costSurfaceToJs(featuresMatrix, output.z);
  redrawGradientPlot(costSurface);

  // run the animated descent
  await runGradientDescent(
    output.z,
    featuresMatrix,
    costSurface,
    maxIterations,
    delay
  );
}

async function runGradientDescent(
  yAxis,
  featuresMatrix,
  costSurface,
  maxIterations,
  delay
) {
  // initialize parameters
  let b = generateRandom();
  let w = Array(featuresMatrix[0][0].length)
            .fill(0)
            .map(() => generateRandom());
  let prevCost = Infinity;
  
  for (let iter = 0; iter < maxIterations; iter++) {
    console.log(maxIterations - iter)
    // 1) compute next step
    const response = gradientDescentToJs(featuresMatrix, yAxis, w, b);
    const { w: newW, b: newB, j: costJ, predictionPlot } = response;

    // 2) update DOM & plots
    spanElemento.textContent = costJ.toFixed(2);
    updatePedictionFunction(newW, newB);

    // add the new surface (prediction plane) &
    addSurfaceToPlot(predictionPlot);

    // plot this point on the cost‐surface
    const [plotW, plotB, plotJ] = adjustGradientPoint(newW[0], newB, costSurface);
    addPointToGradientPlot(plotW, plotB, plotJ, iter > 0);

    // 3) Prepare for next iteration
    w = newW;
    b = newB;

    // stop early if cost isn’t changing
    if (Math.abs(prevCost - costJ) < 1e-6) break;
    prevCost = costJ;

    // 4) wait so user can see it
    await sleep(delay);
  }
}

// simple sleep helper
function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

// (your existing helpers)
function generateRandom() {
  const max = 5, min = -5;
  return Math.random() * (max - min) + min;
}

function roundToNearestPointCostSurface(num) {
  return Math.round(num + 10);
}

function adjustGradientPoint(wAxis, bAxis, costSurface) {
  const i = roundToNearestPointCostSurface(wAxis);
  const j = roundToNearestPointCostSurface(bAxis);
  return [i, j, costSurface[10][10]];
}
