import org.junit.jupiter.api.Test;


import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class AppTest {

    @Test
    void setup() {
    }

    @Test
    void draw() {
    }

    @Test
    void segment() {
    }

    @Test
    void mouseClicked() {
    }

    @Test
    void settings() {
        App app = new App();
        App spy = spy(app);

        spy.main(new String[]{});

        verify(spy).settings();
        verify(spy).size(640, 360);
    }

    @Test
    void main() {
    }
}