import { drawPlots } from './plots/plots.js';
import { start } from './plots/start.js';

const goButton = document.getElementById("btn");

const go = new Go();

			WebAssembly.instantiateStreaming(fetch("assets/main.wasm"), go.importObject).then((result) => {

				go.run(result.instance);

        goButton.addEventListener("click",()=>{
          let input =  document.getElementById("modeloLinear").value;
          try {
            start(input)
        } catch (err) {
          alert(err.message);
        }
        })
			});
      
drawPlots();
