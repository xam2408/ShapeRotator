package main

import (
    "encoding/json"
    "math"
    "time"
    "log"
    "net/http"
    "html/template"
)

type Point3D struct{
    X, Y, Z float64
}

type Screen struct{
    X float64 `json:"x"`
    Y float64 `json:"y"`
    Color string `json:"color"`
}

func getCube() []Point3D{
    return []Point3D{

        { -1, -1, -1 },
        {  1, -1, -1 },
        {  1,  1, -1 },
        { -1,  1, -1 },

        { -1, -1,  1 },
        {  1, -1,  1 },
        {  1,  1,  1 },
        { -1,  1,  1 },

    }
}


func rotateY(p Point3D, angle float64) Point3D{
    nx := p.X*math.Cos(angle) - p.Z*math.Sin(angle) 
    nz := p.X*math.Sin(angle) + p.Z*math.Cos(angle) 
    return Point3D{
        X: nx,
        Y: p.Y,
        Z: nz,
    }    
}

func rotateX(p Point3D, angle float64) Point3D{
    ny := p.Y*math.Cos(angle) - p.Z*math.Sin(angle) 
    nz := p.Y*math.Sin(angle) + p.Z*math.Cos(angle) 
    return Point3D{
        X: p.X,
        Y: ny,
        Z: nz,
    }    
     
}

func proj(p Point3D) Screen{
    scale := 300.0
    distance := 4.0
    center := 300.0

    ooz := 1.0 / (distance - p.Z)

    xp := (p.X * ooz * scale) + center
    yp := (p.Y * ooz * scale) + center

    return Screen{
        X: xp,
        Y: yp,
        Color: "lime",
    }
}


func prepareCube(angleX, angleY float64) []Screen{
    rawPoints := getCube()
    var result []Screen

    for _, p := range rawPoints{
        pRotatedX := rotateX(p, angleX)
        pRotatedY := rotateY(pRotatedX, angleY)

        p2D := proj(pRotatedY)
        result = append(result, p2D)
    }

    return result
}


func indexHandler(w http.ResponseWriter, r *http.Request){
    tmpl, err := template.ParseFiles("index.html")
    if err != nil{
        http.Error(w, "HTML file loading ERROR", 500)
        log.Println(err)
        return
    }
        
        //t := time.Now().UnixMilli()
        //angle := float64(t)/1000.0
        //readyFrame := prepareCube(angle)

        tmpl.Execute(w, nil)
    
}

func dataHandler(w http.ResponseWriter, r *http.Request){
    t := time.Now().UnixMilli()
    angleY := float64(t)/1000.0
    angleX := float64(t)/1500.0

    readyFrame := prepareCube(angleX, angleY)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(readyFrame)

}

func main(){
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/data", dataHandler)
    http.ListenAndServe(":8080", nil)
}


