class Landscape {

  float yOff;

  Landscape(float y) {
    yOff=y;
  }

  void drawHills() {    
    fill(222);// We are going to draw a polygon out of the wave points
    noStroke();
    beginShape();   

    // Iterate over horizontal pixels
    for (float x = 0; x <= width; x += 10) {      
      float y =  hillLevel(x);
      // Set the vertex
     
      vertex(x, y); 
      
    }    
    vertex(width, height);
    vertex(0, height);
    endShape(CLOSE);
  }
  
  void drawGround() {    
    fill(255);// We are going to draw a polygon out of the wave points
    noStroke();
    beginShape();   

    // Iterate over horizontal pixels
    for (float x = 0; x <= width; x += 10) {      
      float y =  groundLevel(x);
      // Set the vertex     
      vertex(x, y);       
    }    
    vertex(width, height);
    vertex(0, height);
    endShape(CLOSE);
  }
   
  
  float hillLevel(float x){
   return map(noise(x*0.02), 0, 1,yOff-height/3,yOff-height/2);  
  }
  
  float groundLevel(float x){
   return map(noise(x*0.005), 0, 1,yOff-20,yOff+20);  
  }
}
