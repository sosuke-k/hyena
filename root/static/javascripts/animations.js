
      init = function(){
        var keys = {38:-0.01, 40:0.01};
        
        var drag_controller = new ControllerDrag('#button',100,200);
        var key_controller = new ControllerKeyboard(keys, true, true);
        var wheel_controller = new ControllerMousewheel(0.01, true);
        
        jarallax = new Jarallax([drag_controller, key_controller, wheel_controller]);
        
        //defaults
        jarallax.setDefault('h1, h2, #p1, #p2, #p3', {display:'none'});
        jarallax.setDefault('.arrow, .question, h1, h2, #p1, #p2, #p3', {opacity:'0'});
        
        //title
        jarallax.addAnimation("h1", [{progress:"0%", opacity:"0"},
                                     {progress:"5%", opacity:"1"},
                                     {progress:"15%", opacity:"1"},
                                     {progress:"20%", opacity:"0"}]);
        jarallax.addAnimation("h1", [{progress:"0%", marginLeft:"100px",  display:"block"}, 
                                     {progress:"20%", marginLeft:"0px"}]);
        
        //slide1
        jarallax.addAnimation(".arrow", [{progress:"20%", top:"-20px"}, {progress:"50%", top:"1px"}]);
        jarallax.addAnimation(".arrow", [{progress:"20%", opacity:"0"}, 
                                         {progress:"30%", opacity:"1"},
                                         {progress:"40%", opacity:"1"},
                                         {progress:"50%", opacity:"0"}]);
        
        jarallax.addAnimation("#head1, #p1", [{progress:"20%", display:"block", marginTop:'20px'}, {progress:"50%", marginTop:'30px'}]);
        jarallax.addAnimation("#head1, #p1", [{progress:"25%", opacity:"0"}, 
                                              {progress:"30%", opacity:"1"},
                                              {progress:"45%", opacity:"1"},
                                              {progress:"50%", opacity:"0"}]);
                                              
        //slide2
        jarallax.addAnimation(".question", [{progress:"50%", marginLeft:"400px"}, {progress:"80%", marginLeft:"380px"}]);
        jarallax.addAnimation(".question", [{progress:"50%", opacity:"0"}, 
                                         {progress:"60%", opacity:"1"},
                                         {progress:"70%", opacity:"1"},
                                         {progress:"80%", opacity:"0"}]);
                                         
        jarallax.addAnimation("#head2, #p2", [{progress:"50%", display:"block", marginLeft:'10px'}, {progress:"80%", marginLeft:'0px'}]);
        jarallax.addAnimation("#head2, #p2", [{progress:"55%", opacity:"0"}, 
                                              {progress:"60%", opacity:"1"},
                                              {progress:"75%", opacity:"1"},
                                              {progress:"80%", opacity:"0"}]);
         
        //slide3                                 
        jarallax.addAnimation("#head3, #p3", [{progress:"80%", display:"block", marginLeft:'0px'}, {progress:"100%", marginLeft:'10px'}]);
        jarallax.addAnimation("#head3, #p3", [{progress:"80%", opacity:"0"}, 
                                              {progress:"90%", opacity:"1"},
                                              {progress:"100%", opacity:"1"}]);
        
      } 

