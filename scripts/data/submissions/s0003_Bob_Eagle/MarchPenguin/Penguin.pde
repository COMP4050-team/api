class Penguin {
  float x, y, dx;
  float target_velocity = 0;
  int SIZE = 50;
  boolean isSelected =  false;

  Beak bill;

  //Constructor to set the initial position, and the intial and target velocity.
  Penguin(float x, float y, float dx) {
    this.x = x;
    this.y = y;
    this.target_velocity = dx;
    this.dx = dx;
    this.bill = new Beak(SIZE*0.45, 0, SIZE/2);
  }


  void display() {
    pushMatrix();
    translate(this.x,this.y);    
    //Black part
    fill(0);
    stroke(1);
    strokeWeight(2);
    ellipse(0, 0, SIZE, SIZE);
    ellipse(0,SIZE, SIZE, 2*SIZE);
    rect(-SIZE/2,SIZE, SIZE, SIZE);
    triangle(0, 0.75*SIZE, 0, 2*SIZE, -SIZE, 2*SIZE);
  
    //The beak
    bill.display();
    
    //The white part
    fill(255);
    noStroke();
    arc(0, 0, SIZE, SIZE, -QUARTER_PI, QUARTER_PI, OPEN);
    arc(SIZE*1.44, 0, 3*SIZE, SIZE, 3*QUARTER_PI, 5*QUARTER_PI, OPEN);
    arc(0, SIZE, SIZE, 2*SIZE, -QUARTER_PI, QUARTER_PI);
    
    //The eyes
    fill(100, 100, 200);
    ellipse(SIZE/4, -SIZE/5, SIZE/6, SIZE/6);
    popMatrix();
  }

  void walk(Landscape land, Penguin[] peers) {
    if (!isSelected) {
      this.x += this.dx;
      this.y = land.groundLevel(this.x)-2*SIZE;
      if (this.x>=width+2*SIZE) {
        this.x = -2*SIZE;
      }
      if (isClose(peers)) {
        this.dx = this.dx-0.05;
      } else {
        this.dx = (9*this.dx+ this.target_velocity)/10;
      }
    }
    this.bill.move();
  }

  void select(float x, float y) {
    this.isSelected=isOver(x, y);
  }

  boolean deSelect() {
    this.isSelected = false;
    return true;
  }

  void drag(float x, float y) {
    if (this.isSelected) {
      this.x = x;
      this.y = y;
      this.bill.squeak();
    }
  }

  boolean isOver(float x, float y) {    
    return dist(x, y, this.x, this.y)<=SIZE;
  }

  boolean isClose(Penguin[] peers) {
    boolean result = false;
    for (int i = 0; i<peers.length; i++) {
      if (this != peers[i]
        && dist(peers[i].x, peers[i].y, this.x, this.y)<=SIZE
        && peers[i].x>this.x) {
        result =true;
      }
    }
    return result;
  }
}
