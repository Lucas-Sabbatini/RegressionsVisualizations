import { drawPlots, redrawPlot} from './plots/plots.js';

const goButton = document.getElementById("btn");

const go = new Go();

			WebAssembly.instantiateStreaming(fetch("assets/main.wasm"), go.importObject).then((result) => {

				go.run(result.instance);

        goButton.addEventListener("click",()=>{
          let input =  document.getElementById("modeloLinear").value;
          try {
          console.log("chamou")
          let output = generateRandomDots(input);
          redrawPlot(output.x,output.y,output.z)
          console.log(output)
        } catch (err) {
          alert("Erro ao executar função: " + err.message);
        }
        })
			});
      
drawPlots();
