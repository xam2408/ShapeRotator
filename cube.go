package main

import (
    "strconv"
    "encoding/json"
    "math"
    "time"
    "log"
    "net/http"
    "html/template"
)

type Point4D struct{
    X, Y, Z, W float64
}

type Point3D struct{
    X, Y, Z float64
}

type Screen struct{
    X float64 `json:"x"`
    Y float64 `json:"y"`
    Color string `json:"color"`
}

func getTessaract() []Point4D{
    var points []Point4D
    for _, x := range []float64{-1,1}{
        for _, y := range []float64{-1,1}{
            for _, z := range []float64{-1,1}{
                for _, w := range []float64{-1,1}{
                    points = append(points, Point4D{x, y, z, w})
                }
            }
        }
    }
    return points
}



func rotateZW(p Point4D, angle float64) Point4D{
    nz := p.Z*math.Cos(angle) - p.W*math.Sin(angle) 
    nw := p.Z*math.Sin(angle) + p.W*math.Cos(angle) 
    return Point4D{
        X: p.X,
        Y: p.Y,
        Z: nz,
        W: nw,
    }
}

func rotateXW(p Point4D, angle float64) Point4D{
    nx := p.X*math.Cos(angle) - p.W*math.Sin(angle) 
    nw := p.X*math.Sin(angle) + p.W*math.Cos(angle) 
    return Point4D{
        X: nx,
        Y: p.Y,
        Z: p.Z,
        W: nw,
    }    
}

func rotateXY(p Point3D, angle float64) Point3D{
    nx := p.X*math.Cos(angle) - p.Y*math.Sin(angle) 
    ny := p.X*math.Sin(angle) + p.Y*math.Cos(angle) 
    return Point3D{
        X: nx,
        Y: ny,
        Z: p.Z,
    }    
     
}

func rotateXZ(p Point3D, angle float64) Point3D{
    nx := p.X*math.Cos(angle) - p.Z*math.Sin(angle) 
    nz := p.X*math.Sin(angle) + p.Z*math.Cos(angle) 
    return Point3D{
        X: nx,
        Y: p.Y,
        Z: nz,
    }    
     
}

func proj4Dto3D(p Point4D) Point3D{
    distance := 2.5
    wFactor := 1.0 / (distance - p.W)
    return Point3D{
        X: p.X * wFactor,
        Y: p.Y * wFactor,
        Z: p.Z * wFactor,

    }
}

func proj(p Point3D) Screen{
    scale := 400.0
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


func prepareTessaract(angle4D, angle3D_Z, angle3D_Y float64) []Screen{
    rawPoints := getTessaract()
    var result []Screen

    for _, p4 := range rawPoints{
        p := rotateZW(p4, angle4D)
   
        p3 := proj4Dto3D(p)
        p3 = rotateXY(p3, math.Pi/4)
        p3 = rotateXZ(p3, math.Pi/4)
        
        p3 = rotateXZ(p3, angle3D_Y)
        
        p2 := proj(p3)
    
        result = append(result, p2)
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
    tStr := r.URL.Query().Get("t")
    var t float64
    if tStr != "" {
        val, err := strconv.ParseFloat(tStr, 64)
        if err == nil{
            t = val
        }
    } else{
            t = float64(time.Now().UnixMilli())
    }

    angle4D := float64(t)/5000.0
    angleZ := float64(t)/3000.0
    angleY := float64(t)/4000.0

    readyFrame := prepareTessaract(angle4D, angleZ, angleY)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(readyFrame)

}

func main(){
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/data", dataHandler)
    http.ListenAndServe(":8080", nil)
}


