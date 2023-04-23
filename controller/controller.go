package controller

import (
    "bytes"
    "encoding/json"
    "fmt"
    "os"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

type ChatCompletion struct {
    Choices []struct {
        FinishReason string `json:"finish_reason"`
        Index        int    `json:"index"`
        Message      struct {
            Content string `json:"content"`
            Role    string `json:"role"`
        } `json:"message"`
    } `json:"choices"`
    Created int64  `json:"created"`
    ID      string `json:"id"`
    Model   string `json:"model"`
    Object  string `json:"object"`
    Usage   struct {
        CompletionTokens int `json:"completion_tokens"`
        PromptTokens     int `json:"prompt_tokens"`
        TotalTokens      int `json:"total_tokens"`
    } `json:"usage"`
}

var TurboApiUrl = "https://api.openai.com/v1/chat/completions"

func SearchFunctions(c *gin.Context) {

    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    data := CreateSearchData("c", "sizeof", []string{"python", "java"})

    payload, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request payload"})
        return
    }

    req, err := http.NewRequest("POST", TurboApiUrl, bytes.NewBuffer(payload))
    if err != nil {
        fmt.Println("Error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create HTTP request"})
        return
    }

    req.Header.Set("Authorization", "Bearer " + os.Getenv("API_KEY"))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send HTTP request"})
        return
    }
    defer resp.Body.Close()

    fmt.Println("Status code:", resp.StatusCode)

    // parse response body
    var responseBody map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&responseBody)
    if err != nil {
        fmt.Println("Error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response body"})
        return
    }

    // extract message from response body
    message, ok := responseBody["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"]
    if !ok {
        fmt.Println("Error: Failed to extract message from response body")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract message from response body"})
        return
    }

    // unmarshal message string to JSON object
    var messageJSON interface{}
    if err := json.Unmarshal([]byte(message.(string)), &messageJSON); err != nil {
        fmt.Println("Error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal message to JSON"})
        return
    }

    // return HTTP response
    c.JSON(http.StatusOK, gin.H{"message": messageJSON})
}

func CreateSearchData(choiceLanguage string, searchFunction string, responseLanguages []string) map[string]interface{} {
    messageDataArray := []map[string]interface{}{
        {
            "role":    "user",
            "content": `{"language": "python", "function": "print", "response": ["python", "java"]}`,
        },
        {
            "role": "assistant",
            "content": `{"python":{"function":"print","args":"可変長文字列","return":"なし","example":"print(\"Hello, World!\") # Output: Hello, World!"},"java":{"function":"System.out.println","args":"可変長文字列","return":"なし","example":"System.out.println(\"Hello, World!\"); // Output: Hello, World!"}}`,
        },
        {
            "role":    "user",
            "content": `{"language": "python", "function": "for", "response": ["python", "java"]}`,
        },
        {
            "role": "assistant",
            "content": `{"python":{"function":"for","args":"オブジェクト","return":"なし","example":"for x in iterable:\n    print(x)"},"java":{"function":"for","args":"初期化式; 条件式; 変化式;","return":"なし","example":"for (int i = 0; i < 10; i++) {\n    System.out.println(i);\n}"}}`,
        },
        {
            "role":    "user",
            "content": `{"language": "` + choiceLanguage + `", "function": "` + searchFunction + `", "response": ` + fmt.Sprintf("%v", responseLanguages) + `}`,
        },
    }

    requestMessage := map[string]interface{}{
        "model":       "gpt-3.5-turbo",
        "temperature": 0.0,
        "messages":     messageDataArray,
    }

    return requestMessage
}
