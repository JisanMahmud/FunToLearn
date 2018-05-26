import React from 'react';

const CndtlRendering = (props)=>{
    if(props.value){
        
        console.log(props)
        console.log(props.value)
        return <td>{props.value}</td>
    }
    else{
        console.log(props)
        return ""
    }
}

export default CndtlRendering