function scrollToEnd(smooth = false){
    let scroll = document.body.clientHeight;
    
    if(smooth == true){
        window.scroll({
            top: scroll,
            left: 0,
            behavior: 'smooth'
        });
    }
    window.scrollTo(0, scroll)
}


window.addEventListener("scroll", function(e){
    console.log(window.scrollY)
})
