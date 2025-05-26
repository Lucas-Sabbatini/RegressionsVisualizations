# Regression Visualizer

I built an interactive web application for visualizing polynomial regression in 3D space—because landing a junior developer role in 2025 is tougher than ever.
This app features real-time gradient descent optimization and beautiful mathematical visualizations.

All implementation was written from scratch in Go, without using any machine-learning libraries.

## Features

- Interactive 3D polynomial regression visualization
- Real-time gradient descent optimization
- Support for custom polynomial terms
- Live cost function visualization
- Mathematical formula rendering using KaTeX
- Beautiful and intuitive user interface

## Technical Details

The application uses several technologies to provide a rich interactive experience:

- **WebAssembly and Go**: For efficient computation of regression parameters
- **Frontend**: Pure HTML, CSS, and JavaScript
- **Mathematical Rendering**: KaTeX for beautiful formula display
- **Plotting**: Plotly.js for interactive 3D visualizations

## Getting Started

### This application is hosted on [GitHub Pages](https://github.com/Lucas-Sabbatini/RegressionsVisualizations)!

### OR...

### Prerequisites

- A modern web browser (Chrome, Firefox, Safari, or Edge recommended)
- No additional installation required - the application runs entirely in the browser

### Usage

1. Open `index.html` in your web browser
2. Enter your desired polynomial terms in the input field using the format:
   - For simple terms: `{x1, x2}`
   - For squared terms: `{x1^2, x2^2}`
   - For mixed terms: `{x1^2, x2^2, x1x2}`
3. Click the "Go!" button to start the visualization
4. Watch as the gradient descent algorithm optimizes the polynomial regression
5. Explore the interactive 3D plots and observe the cost function minimization

## Project Structure

```
regressionVisualizer/
├── assets/           # Static assets including WebAssembly files
├── plots/           # Plot generation and visualization code
├── src/             # Source code for regression calculations
├── index.html       # Main application page
├── style.css        # Styling and layout
├── main.js          # Application entry point
└── formatting.js    # Mathematical formatting utilities
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- Thanks to all the open-source libraries that made this project possible
- Special thanks to the Plotly.js and KaTeX teams for their excellent visualization and math rendering tools 
