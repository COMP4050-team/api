/* A march of the Penguins animation
   Created for Comp4050
   2022
   Carols's submission
*/

Penguin[] colony = new Penguin[7];
Landscape southPole = new Landscape(500);
void setup() {
  size(1000, 600);

  for (int i =0; i<colony.length; i++) {
    colony[i]= new Penguin(random(0, width), 400,random(0.5,1.5));
  }
}


void draw() {
  background(200, 0, 100);
  southPole.drawHills();
  for (int i = 0; i<colony.length; i++) {
    colony[i].display();
    colony[i].walk(southPole,colony);
  }
  southPole.drawGround();
}

void mousePressed(){
  for (int i = 0; i<colony.length; i++) {
    colony[i].select(mouseX,mouseY);
  }
}

void mouseReleased(){
  for (int i = 0; i<colony.length; i++) {
    colony[i].deSelect();
  }
}

void mouseDragged(){
  for (int i = 0; i<colony.length; i++) {
    colony[i].drag(mouseX,mouseY);
  }
}
