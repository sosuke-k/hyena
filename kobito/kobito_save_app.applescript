f u n c t i o n   k o b i t o _ g e t _ w i n d o w _ i n f o ( ) { 
 	 v a r   s y s t e m E v e n t   =   A p p l i c a t i o n ( ' S y s t e m   E v e n t s ' ) ; 
 	 s y s t e m E v e n t . i n c l u d e S t a n d a r d A d d i t i o n s   =   t r u e ; 
 
 	 v a r   k o b i t o P r o c e s s   =   s y s t e m E v e n t . p r o c e s s e s . b y N a m e ( ' K o b i t o ' ) ; 
 	 / /   I f   t h e   s y s t e m   l a n g u a g e   i s   E n g l i s h 
 	 v a r   w i n d o w M e n u   =   k o b i t o P r o c e s s . m e n u B a r s [ 0 ] . m e n u B a r I t e m s . b y N a m e ( ' W i n d o w ' ) ; 
 	 / /   I f   t h e   s y s t e m   l a n g u a g e   i s   J a p a n e s e 
 	 / /   w i n d o w M e n u   =   k o b i t o P r o c e s s . m e n u B a r s [ 0 ] . m e n u B a r I t e m s . b y N a m e ( ' �0�0�0�0�0' ) ; 
 	 v a r   w i n d o w I t e m s   =   w i n d o w M e n u . m e n u s [ 0 ] . m e n u I t e m s ; 
 	 v a r   r e s   =   { } ; 
 	 r e s [ 0 ]   =   [ ] ; 
 	 
 	 v a r   i   =   w i n d o w I t e m s . l e n g t h   -   1 ; 
 	 w h i l e   ( i   >   0 ) { 
 	 	 i f   ( w i n d o w I t e m s [ i ] . n a m e ( )   = =   n u l l ) { 
 	 	 	 b r e a k ; 
 	 	 } 
 	 	 r e s [ 0 ] . p u s h ( w i n d o w I t e m s [ i ] . n a m e ( ) ) ; 
 	 	 i - - ; 
 	 } 
 	 
         v a r   d a t a   =   J S O N . s t r i n g i f y ( r e s ) ; 
         r e t u r n   d a t a ; 
 } 
 
 
 f u n c t i o n   f i l e W r i t e r ( p a t h A s S t r i n g )   { 
         ' u s e   s t r i c t ' ; 
 
         v a r   a p p   =   A p p l i c a t i o n . c u r r e n t A p p l i c a t i o n ( ) ; 
         a p p . i n c l u d e S t a n d a r d A d d i t i o n s   =   t r u e ; 
         v a r   p a t h   =   P a t h ( p a t h A s S t r i n g ) ; 
         v a r   f i l e   =   a p p . o p e n F o r A c c e s s ( p a t h ,   { 
                 w r i t e P e r m i s s i o n :   t r u e 
         } ) ; 
 
         / *   r e s e t   f i l e   l e n g t h   * / 
         a p p . s e t E o f ( f i l e ,   { 
                 t o :   0 
         } ) ; 
 
         r e t u r n   { 
                 w r i t e :   f u n c t i o n ( c o n t e n t )   { 
                         a p p . w r i t e ( c o n t e n t ,   { 
                                 t o :   f i l e , 
                                 a s :   ' t e x t ' 
                         } ) ; 
                 } , 
                 c l o s e :   f u n c t i o n ( )   { 
                         a p p . c l o s e A c c e s s ( f i l e ) ; 
                 } 
         } ; 
 } 
 
 
 f u n c t i o n   r u n ( a r g v ) { 
 	 v a r   a c t i v e W i n d o w s   =   k o b i t o _ g e t _ w i n d o w _ i n f o ( ) ; 
 	 c o n s o l e . l o g ( a c t i v e W i n d o w s ) ; 	 
 	 v a r   e x p o r t F i l e W r i t e r   =   f i l e W r i t e r ( a r g v ) ; 
 	 e x p o r t F i l e W r i t e r . w r i t e ( a c t i v e W i n d o w s ) ; 
 	 c o n s o l e . l o g ( " w r i t t e n " ) ; 
 	 e x p o r t F i l e W r i t e r . c l o s e ( ) ; 
 } 
