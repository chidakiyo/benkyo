import React, { useEffect, useState } from 'react'

const useLocalStorage = (key, defaultValue) => {
    const [value, setValue] = useState(() => {
        const jsonValue = window.localStorage.getItem(key);
        if (jsonValue !== null) return JSON.parse(jsonValue)
        return defaultValue
    });

    // 発火のタイミングを決めることができる
    useEffect(() => {
        console.log("### localstorageに書き込む " + value)
        window.localStorage.setItem(key, JSON.stringify(value));
    }, [value, setValue])

    return [value, setValue]
};

export default useLocalStorage