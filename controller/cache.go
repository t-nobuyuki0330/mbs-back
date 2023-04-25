package controller

import (
    "database/sql"
    "encoding/json"
    "errors"
    "fmt"
    "math/rand"
    "time"
)

func SelectCache( db *sql.DB, req_lang string, req_func string, req_resp string ) ( map[string]interface{}, error ) {
    rows, err := db.Query( fmt.Sprintf( "SELECT ans_json::JSON FROM public.cache WHERE req_lang = '%s' AND req_func = '%s' AND req_resp = '%s' AND ans_json IS NOT NULL", req_lang, req_func, req_resp ) )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    results := []map[string]interface{}{}
    for rows.Next() {
        var resultJSON []byte
        if err := rows.Scan( &resultJSON ); err != nil {
            return nil, err
        }

        var result map[string]interface{}
        if err := json.Unmarshal( resultJSON, &result ); err != nil {
            return nil, err
        }

        results = append( results, result )
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    if len( results ) == 0 {
        return nil, errors.New( "no results found" )
    }

    // シード値を現在時刻に設定する
    rand.Seed( time.Now().UnixNano() )
    // スライスのインデックスをランダムに選択する
    randomIndex := rand.Intn( len( results ) )
    // ランダムに選択した値を返す
    return results[randomIndex], nil
}
