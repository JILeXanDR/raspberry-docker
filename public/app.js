const ws = new ReconnectingWebSocket(`ws://${window.location.hostname}:8000/ws`);
const touchButton = document.querySelector('#touch_button_state');
const echoLocator = document.querySelector('#echo_locator_distance');
const audio = new Audio('/beep.mp3');

const setState = (text, color) => {
  touchButton.textContent = text;
  touchButton.style.color = color;
};

setState('undefined', '#000000');

ws.addEventListener('open', () => {
  ws.send('ping');
});

ws.addEventListener('message', (event) => {
  const value = event.data;
  if (value === 'on') {
    setState('pressed', '#0cf811');
  } else if (value === 'off') {
    setState('not pressed', '#f83b60');
  } else if (value.includes('centimeters')) {
    echoLocator.textContent = value;
    const centimeters = parseFloat(value.replace(' centimeters'));
    if (centimeters <= 70) {
      audio.play();
    }
  } else {
    console.warn('unhandled message: %v', value);
  }
});
