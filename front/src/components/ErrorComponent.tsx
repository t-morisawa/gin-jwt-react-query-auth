import { useEffect } from "react";
import { storage } from "../utils";

export default function ErrorComponent(error) {
    useEffect(() => {
        storage.clearToken();
    }, []);

    return (
        <div style={{color: 'tomato'}}>
            {JSON.stringify(error, null, 2)}
        </div>
    );
  }
  