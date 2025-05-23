import { drawPlots} from './plots/plots.js';

const goButton = document.getElementById("btn");

const go = new Go();

			WebAssembly.instantiateStreaming(fetch("assets/main.wasm"), go.importObject).then((result) => {

				go.run(result.instance);

        goButton.addEventListener("click",()=>{
          let input =  document.getElementById("modeloLinear").value;
          try {
          generateRandomDots(input);
        } catch (err) {
          alert("Erro ao executar função: " + err.message);
        }
        })
			});
      
drawPlots();
