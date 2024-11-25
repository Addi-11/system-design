## Implement HLS Video Streaming

Run the following command where the .mp4 file is saved.
```
ffmpeg -i song.mp4 -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list outputlist.m3u8 -segment_format mpegts output%03d.ts
```

Output from port `http://localhost:8080/outputlist.m3u8`

```
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-MEDIA-SEQUENCE:0
#EXT-X-ALLOW-CACHE:YES
#EXT-X-TARGETDURATION:11
#EXTINF:10.066667,
output000.ts
#EXTINF:10.300000,
output001.ts
#EXTINF:9.766667,
output002.ts
#EXTINF:5.333333,
output003.ts
#EXT-X-ENDLIST
```

Reference: [Video Streaming in GO](https://medium.com/bootdotdev/create-a-golang-video-streaming-server-using-hls-a-tutorial-f8c7d4545a0f)