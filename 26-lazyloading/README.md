## Lazy Loading Images

Initially, only a few images are loaded, while the remaining images are loaded as you scroll down the page, reducing initial load time.

### Key parts enabling Lazy Loading

#### 1. IntersectionObserver:
- The `IntersectionObserver` watches each image to see if it’s intersecting with the viewport. When an image starts to enter the viewport, the observer triggers the lazyLoad callback function.

#### 2. Data Attribute for Image Source (data-src):
- Each <img> tag has  `data-src` attribute containing the image’s actual source, not src. This prevents the browser from loading the image immediately.
- Inside the lazyLoad function, the code sets `img.src = img.dataset.src` to load the image only when it’s in view.

#### 3. Lazy Loading Class (lazy):
- Once an image enters the viewport and the actual src is set, and the lazy-loaded class is added to make it visible with a smooth fade-in.
- `rootMargin`: Controls when to start loading by extending the viewport bounds. 
- `threshold`: Sets how much of the image needs to be visible to trigger the loading, allowing partial visibility to load the image. <br><br>

Before Scrolling
![](../images/lazyloading.png)

After Scrolling
![](../images/lazyloading2.png)