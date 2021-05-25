clockd1_={
    "indicate": true,
    "indicate_color": "#FFFFFF",
    "dial1_color": "#FF1493",
    "dial2_color": "#EE1289",
    "dial3_color": "#CD1076",
    "time_add": 1,
    "time_24h": true,
    "time_add_color": "#FF4500",
    "date_add": 3,
    "date_add_color": "#FFD700",
    "bg_color": "#69696930",
   };

var c = document.getElementById('clock');
cns1_ = c.getContext('2d');

clock_follow(350,cns1_,clockd1_);
