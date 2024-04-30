import React, { useRef, useEffect } from 'react';

const EbitenGame = () => {
  const canvasRef = useRef(null);

  useEffect(() => {
    const initializeGame = () => {
      if (!window.Go) {
        console.error("Ebiten Go library (window.Go) not found.");
        return;
      }
      
      const go = new window.Go();

      const runGame = async () => {
        const wasmPath = './main.wasm'; // Update with the path to your main.wasm
        const response = await fetch(wasmPath);
        const buffer = await response.arrayBuffer();
        const { instance } = await WebAssembly.instantiate(buffer, go.importObject);
        go.run(instance);
      };

      runGame();
    };

    // Ensure the DOM content is loaded before initializing the game
    document.addEventListener('DOMContentLoaded', initializeGame);

    return () => {
      // Cleanup if needed
      document.removeEventListener('DOMContentLoaded', initializeGame);
    };
  }, []);

  return <canvas ref={canvasRef} id="gameCanvas" width="640" height="480"></canvas>;
};

export default EbitenGame;
