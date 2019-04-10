const app = new Vue({
  el: '#app',
  data: {
    touchButtonState: 'undefined',
    echoLocatorDistance: 'undefined',
  },
  created() {
    const ws = new ReconnectingWebSocket(`ws://${window.location.hostname}:8000/ws`);
    const audio = new Audio('/beep.mp3');

    ws.addEventListener('open', () => {
      ws.send('ping');
    });

    ws.addEventListener('message', (event) => {
      const value = event.data;
      if (value === 'on') {
        this.touchButtonState = 'pressed';
      } else if (value === 'off') {
        this.touchButtonState = 'not pressed';
      } else if (value.includes('centimeters')) {
        this.echoLocatorDistance = value;
        const centimeters = parseFloat(value.replace(' centimeters'));
        if (centimeters <= 70) {
          audio.play();
        }
      } else {
        console.warn('unhandled message: %v', value);
      }
    });
  },
});
