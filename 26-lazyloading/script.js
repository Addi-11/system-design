document.addEventListener("DOMContentLoaded", function() {
    const lazyImages = document.querySelectorAll("img.lazy");

    const lazyLoad = (entries, observer) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                const img = entry.target;

                // adding delay before loading image
                setTimeout(() => {
                    img.src = img.dataset.src; // load image
                    img.onload = () => {
                        img.classList.add("lazy-loaded");
                    };
                }, 3000); // add delay - to simulate network delay

                observer.unobserve(img); // stop observing once loaded
            }
        });
    };

    const observer = new IntersectionObserver(lazyLoad, {
        root: null, // viewport
        rootMargin: "0px 0px 20px 0px", //  offset that expands or contracts the viewport by the specified values to determine if an element is in view.
        threshold: 0.3 // proportion of the target’s visibility needed to trigger the callback.
    });

    lazyImages.forEach(img => observer.observe(img));
});
