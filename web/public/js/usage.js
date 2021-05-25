week_arr=["Sunday", "Monday", "Tuesday","Wednesday","Thursday","Friday","Saturday"];
month_arr=["January","February","March","April","May","June","July","August","September","October","November","December"];

var c=document.getElementById("usage");
var ctx=c.getContext("2d");
ctx.fillStyle="white";
ctx.strokeStyle="white";
ctx.beginPath();
ctx.lineWidth=3;
ctx.moveTo(0+3,350/2);
ctx.lineTo(900,350/2);
ctx.stroke();

ctx.beginPath();
ctx.lineWidth=2;
ctx.arc(-350/2+20,350/2,350/2,(-0.1 * Math.PI), (0.1 * Math.PI));
ctx.stroke();
ctx.closePath();

ctx.beginPath();
ctx.lineWidth=2;
ctx.arc(-350/2+430,350/2,350/2,(-0.1 * Math.PI), (0.1 * Math.PI));
ctx.stroke();
ctx.closePath();

ctx.fillStyle="#D6A29C";
ctx.strokeStyle="#D6A29C";

var now=new Date();

var sec=now.getSeconds();
var min=now.getMinutes();
var hour=now.getHours();

if(hour==0){hour=12;}
if(hour<10){hour="0"+hour;}
if(min<10){min="0"+min;}
if(sec<10){sec="0"+sec;}

var day=now.getDate();
var week=now.getDay();
var year=now.getFullYear();
var month=now.getMonth()+1;

if((month)<10){month="0"+(month);}
if(day<10){day="0"+day;}

ctx.lineWidth=1;
ctx.textBaseline="middle";
ctx.textAlign="center";
ctx.font=54+"px 'Comic Neue'";

ctx.fillText(hour+":"+min,100,60);

ctx.font=26+"px 'Comic Neue'";
ctx.fillText(":"+sec,180,50);

ctx.font=28+"px 'Amaze'";
week=now.getDay();
month=now.getMonth();
ctx.fillText(week_arr[week]+" - "+month_arr[month] +"  "+day+","+year,300,85);

ctx.fillStyle="#494a5b";
ctx.strokeStyle="#494a5b";

ctx.lineWidth=1;
ctx.textBaseline="middle";
ctx.textAlign="center";
ctx.font=22+"px '迷你简硬笔行书";
ctx.fillText("给岁月以文明,给时光以生命。",430+220,350/2-20);

ctx.lineWidth=1;
ctx.textBaseline="middle";
ctx.textAlign="center";
// ctx.font=22+"px 'Comic Neue'";
ctx.fillText("To the time to life, rather than to life in time.",430+220,350/2+20);