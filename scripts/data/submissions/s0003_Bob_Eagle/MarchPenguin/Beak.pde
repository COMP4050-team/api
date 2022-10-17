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
    fill(200, 150, 100);
    stroke(210, 160, 120);
    pushMatrix();
    translate(x, y);
    rotate(radians(angle));
    triangle(0, 0, 0, size*0.45, size, 0);
    rotate(radians(-2*angle));
    triangle(0, 0, 0, -size*0.45, size, 0);
    popMatrix();
  }

  void move() {
    if (squeakCounter>0) {
      if (angle<=0 && angularSpeed<0) {
        angularSpeed = - angularSpeed;
        squeakCounter--;
      }
      if (angle>=45 && angularSpeed>0) {
        angularSpeed = - angularSpeed;
      }
      angle+=angularSpeed;
    }
  }

  void squeak() {
    squeakCounter= 3;
  }
}
