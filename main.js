const go = new Go();

			WebAssembly.instantiateStreaming(fetch("assets/main.wasm"), go.importObject).then((result) => {

				go.run(result.instance);

        generateRandomDots();
			});
      
const getRandomValues = n =>
  Array.from({ length: n }, () => Math.random() * 200 - 100);

var n = 100
var trace1 = {
	x: getRandomValues(n), 
  y: getRandomValues(n), 
  z: getRandomValues(n),

	mode: 'markers',

	marker: {
		size: 3,
		line: {
		color: 'rgba(217, 217, 217, 0.14)',
		width: 0.5},
		opacity: 0.8},
	type: 'scatter3d'
};

var data = [trace1];
Plotly.newPlot('plot', data);