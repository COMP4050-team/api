class Beak {

  float x, y, size;
  int angle = 0;
  int angularSpeed = 3;
  int squeakCounter = 0;


  Beak(float x, float y, float size) {
    this.x = x;
    this.y = y;
    this.size = size;
  }

  void display() {
    
  }

  void move() {        
      angle+=angularSpeed;
  }

  void squeak() {
    squeakCounter= 3;
  }
}
