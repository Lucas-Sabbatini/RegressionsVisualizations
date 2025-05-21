let pyodideReadyPromise = (async () => {
        const pyodide = await loadPyodide();
        await pyodide.loadPackage('micropip');
        await pyodide.runPythonAsync(`
import micropip
await micropip.install("plotly")
`);
        // ⚠️ Carrega e executa o script externo main.py
        const response = await fetch("main.py");
        const code = await response.text();
        await pyodide.runPythonAsync(code);
        return pyodide;
      })();

      document.getElementById("btn").addEventListener("click", async () => {
        const pyodide = await pyodideReadyPromise;
        const figDict = await pyodide.runPythonAsync("get_plot_dict()");
        Plotly.newPlot('plot', figDict.data, figDict.layout);
      });