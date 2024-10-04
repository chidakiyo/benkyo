import { useEffect, useState, useContext, useRef, useReducer, useMemo, useCallback } from 'preact/hooks';
import './app.css'
import PersonalInfoContext from './info';
import SomeChild from './SomeChild';
import useLocalStorage from './useLocalStorage';

const reducer = (state, action) => {
  switch (action.type) {
    case "increment":
      return state + 1;
    case "decrement":
      return state - 1;
    default:
      return
  }
}

export function App() {

  const [count, setCount] = useState(0);
  const personalInfo = useContext(PersonalInfoContext)
  const ref = useRef()
  const [state, dispatch] = useReducer(reducer, 0)

  const handleClick = () => {
    setCount(count + 1);
  };

  // 発火のタイミングを決めることができる
  // 副作用
  useEffect(() => {

    console.log("hello")

    // からの[]でロードされたタイミング
  }, [count])

  const handleRef = () => {
    console.log(ref.current.value);
  }

  // useMemo
  // メモすることができる : メモ化,,ブラウザのメモリに保存できる
  const [count01, setCount01] = useState(0);
  const [count02, setCount02] = useState(0);

  // const square = () => {
  //   let i = 0;
  //   while (i < 2000000) {
  //     i++;
  //   }
  //   return count02 * count02;
  // }

  const square = useMemo(() => {
    let i = 0;
    // 重い処理
    while (i < 200000000) {
      i++;
    }
    return count02 * count02;
  }, [count02])

  // useCallBack 関数のメモ化
  const [counter, setCounter] = useState(0);

  // const showCount = () => {
  //   alert("これは重い処理です")
  // }

  const showCount = useCallback(() => {
    alert("これは重い処理です")
  }, [counter]);

  // カスタムフック
  const [age, setAge] = useLocalStorage("age", 44)

  return (
    <div className="App">
      <h1>useState, useEffect</h1>
      <button onClick={handleClick}>+</button>
      <p>{count}</p>

      <hr />

      <h1>useContext</h1>
      <p>{personalInfo.name}</p>
      <p>{personalInfo.age}</p>

      <hr />

      <h1>useRef</h1>
      <input type="text" ref={ref} />
      <button onClick={handleRef}>UseRef</button>

      <hr />

      <h1>useReducer</h1>
      <p>カウント : {state}</p>
      <button onClick={() => dispatch({
        type: "increment"
      })}>+</button>
      <button onClick={() => dispatch({
        type: "decrement"
      })}>-</button>

      <hr />

      <h1>useMemo</h1>
      <p>カウント1 : {count01}</p>
      <p>カウント2 : {count02}</p>
      <p>結果 : {square}</p>
      <button onClick={() => setCount01(count01 + 1)}>+</button>
      <button onClick={() => setCount02(count02 + 1)}>+</button>

      <hr />
      <h1>useCallback</h1>
      <SomeChild showCount={showCount} />

      <hr />
      <h1>custom hooks</h1>
      <p>{age}</p>
      <button onClick={() => setAge(80)}>年齢をセット</button>

    </div>
  )
}
