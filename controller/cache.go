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

func RegistCache( db *sql.DB, req_lang string, req_func string, req_resp string ) ( int, error ) {
    // SQLクエリを作成する
    var id int
    query := "INSERT INTO cache ( req_lang, req_func, req_resp, req_count, ans_json ) VALUES ( $1, $2, $3, $4, $5 ) RETURNING id"

    // SQLクエリを実行する
    err := db.QueryRow( query, req_lang, req_func, req_resp, 0, nil ).Scan( &id )
    if err != nil {
        return 0, err
    }

    return id, nil
}

func UpdateCache( db *sql.DB, id int, ans_resp string ) error {
    // JSONを文字列に変換する
    ans_respString, err := json.Marshal( ans_resp )
    if err != nil {
        return err
    }

    // キャッシュレコードを更新する
    result, err := db.Exec( "UPDATE public.cache SET ans_json = $1 WHERE id = $2", ans_respString, id )
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return errors.New( "cache record not found" )
    }

    return nil
}
