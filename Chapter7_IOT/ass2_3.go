a. Block Diagram / Pin Diagram of Arduino Uno Interfacing with an IR Sensor
Block Diagram: Arduino Uno with IR Sensor


Pin Connections:

IR Sensor:
VCC (Power) → 5V on Arduino
GND (Ground) → GND on Arduino
OUT (Signal) → Digital Pin 2 on Arduino

-------------------------------------------------------
C++ Code for Arduino: Turn ON/OFF Buzzer
---------------------------------------------------------
// Define the pin number for the buzzer
const int buzzerPin = 9;

void setup() {
  // Initialize the buzzer pin as an output
  pinMode(buzzerPin, OUTPUT);
}

void loop() {
  // Turn the buzzer on
  digitalWrite(buzzerPin, HIGH);
  delay(1000); // Wait for 1 second
  // Turn the buzzer off
  digitalWrite(buzzerPin, LOW);
  delay(1000); // Wait for 1 second
}

-------------------------------------------------------------------


c. Observations on Input and Output
Observations:

Input: The program does not take any user input; it runs continuously until interrupted.
Output: The buzzer connected to the specified pin turns ON and OFF every second. The audible output can be heard directly from the buzzer.
