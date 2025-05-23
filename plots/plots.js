const getRandomValues = n =>
  Array.from({ length: n }, () => Math.random() * 200 - 100);


export function drawPlots() {
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
    const layout = {
    title: {
        text: 'Dataset and Predictions'
    },
    scene: {
        xaxis: {
        title: {
            text: 'X1'
        }
        },
        yaxis: {
        title: {
            text: 'X2'
        }
        },
        zaxis: {
        title: {
            text: 'Y'
        }
        }
    },
    autosize: false,
    width: 400,
    height: 400,
    margin: {
        l: 10,
        r: 10,
        b: 10,
        t: 40,
    }
    };
    var data = [trace1];
    Plotly.newPlot('plot', data,layout);

    const z1 = [

        [8.83,8.89,8.81,8.87,8.9,8.87],

        [8.89,8.94,8.85,8.94,8.96,8.92],

        [8.84,8.9,8.82,8.92,8.93,8.91],

        [8.79,8.85,8.79,8.9,8.94,8.92],

        [8.79,8.88,8.81,8.9,8.95,8.92],

        [8.8,8.82,8.78,8.91,8.94,8.92],

        [8.75,8.78,8.77,8.91,8.95,8.92],

        [8.8,8.8,8.77,8.91,8.95,8.94],

        [8.74,8.81,8.76,8.93,8.98,8.99],

        [8.89,8.99,8.92,9.1,9.13,9.11],

        [8.97,8.97,8.91,9.09,9.11,9.11],

        [9.04,9.08,9.05,9.25,9.28,9.27],

        [9,9.01,9,9.2,9.23,9.2],

        [8.99,8.99,8.98,9.18,9.2,9.19],

        [8.93,8.97,8.97,9.18,9.2,9.18]

    ];

    var data_z1 = {z: z1, type: 'surface',  showscale: false};

    const layoutGradient = {
    title: {
        text: 'Gradient descent'
    },
    scene: {
        xaxis: {
        title: {
            text: 'W'
        },
        showticklabels: false,
        },
        yaxis: {
        title: {
            text: 'B'
        },
        showticklabels: false,
        },
        zaxis: {
        title: {
            text: 'J'
        },
        showticklabels: false,
        }
    },
    autosize: false,
    width: 400,
    height: 400,
    margin: {
        l: 10,
        r: 10,
        b: 10,
        t: 40,
    }
    };
    Plotly.newPlot('gradientPlot', [data_z1],layoutGradient);
}
