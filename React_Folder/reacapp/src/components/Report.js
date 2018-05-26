import React, { Component } from 'react';
import GetReportXML from './GetRequest';
import CndtlRendering from './ConditionalReturn'

class Report extends Component{
    constructor(props){
        super(props);
    
        var dataTable = [];
        console.log(dataTable)
        this.state={
            'Loaded':false,
            'ThisT' : []
        };

        
    }

    componentDidMount(){
        console.log("Comp mounted")
        this.myMethod();
    }

    myMethod(){
        this.setState({Loaded : true})
        var JSONTable = GetReportXML("/GetReport/1stReport")

        var Table = JSONTable.Rows.map((row)=>{
            return(
                <tr key={row.Field1}>
                    <CndtlRendering value={row.Field1}/>
                    <CndtlRendering value={row.Field2}/>
                    <CndtlRendering value={row.Field3}/>
                    <CndtlRendering value={row.Field4}/>
                    <CndtlRendering value={row.Field5}/>
                    <CndtlRendering value={row.Field6}/>
                    <CndtlRendering value={row.Field7}/>
                    <CndtlRendering value={row.Field8}/>
                    <CndtlRendering value={row.Field9}/>
                    <CndtlRendering value={row.Field10}/>
                    <CndtlRendering value={row.Field11}/>
                    <CndtlRendering value={row.Field12}/>
                    <CndtlRendering value={row.Field13}/>
                    <CndtlRendering value={row.Field14}/>
                    <CndtlRendering value={row.Field16}/>
                    <CndtlRendering value={row.Field17}/>
                    <CndtlRendering value={row.Field18}/>
                    <CndtlRendering value={row.Field19}/>
                    <CndtlRendering value={row.Field20}/>
                    <CndtlRendering value={row.Field21}/>
                    <CndtlRendering value={row.Field22}/>
                    <CndtlRendering value={row.Field23}/>
                    <CndtlRendering value={row.Field24}/>
                    <CndtlRendering value={row.Field25}/>
                    <CndtlRendering value={row.Field26}/>
                    <CndtlRendering value={row.Field27}/>
                    <CndtlRendering value={row.Field28}/>
                    <CndtlRendering value={row.Field29}/>
                    <CndtlRendering value={row.Field30}/>
                    <CndtlRendering value={row.Field31}/>
                    <CndtlRendering value={row.Field32}/>
                    <CndtlRendering value={row.Field33}/>
                    <CndtlRendering value={row.Field34}/>
                    <CndtlRendering value={row.Field35}/>
                    <CndtlRendering value={row.Field36}/>
                    <CndtlRendering value={row.Field37}/>
                    <CndtlRendering value={row.Field38}/>
                    <CndtlRendering value={row.Field39}/>
                    <CndtlRendering value={row.Field40}/>
                    <CndtlRendering value={row.Field41}/>
                    <CndtlRendering value={row.Field42}/>
                    <CndtlRendering value={row.Field43}/>
                    <CndtlRendering value={row.Field44}/>
                    <CndtlRendering value={row.Field45}/>
                    <CndtlRendering value={row.Field46}/>
                    <CndtlRendering value={row.Field47}/>
                    <CndtlRendering value={row.Field48}/>
                    <CndtlRendering value={row.Field49}/>
                    <CndtlRendering value={row.Field50}/>
                    <CndtlRendering value={row.Field51}/>
                    <CndtlRendering value={row.Field52}/>
                    <CndtlRendering value={row.Field53}/>
                    <CndtlRendering value={row.Field54}/>
                    <CndtlRendering value={row.Field55}/>
                    <CndtlRendering value={row.Field56}/>
                    <CndtlRendering value={row.Field57}/>
                    <CndtlRendering value={row.Field58}/>
                    <CndtlRendering value={row.Field59}/>
                    <CndtlRendering value={row.Field60}/>
                    <CndtlRendering value={row.Field61}/>
                    <CndtlRendering value={row.Field62}/>
                    <CndtlRendering value={row.Field63}/>
                    <CndtlRendering value={row.Field64}/>
                    <CndtlRendering value={row.Field65}/>
                    <CndtlRendering value={row.Field66}/>
                    <CndtlRendering value={row.Field67}/>
                    <CndtlRendering value={row.Field68}/>
                    <CndtlRendering value={row.Field69}/>
                    <CndtlRendering value={row.Field70}/>
                </tr>); 
        })

        this.dataTable = Table 
        console.log(this.dataTable)
    }

    render(){
        return(
            <div>
                <table>
                    <tbody>
                        {this.dataTable}
                    </tbody>
                </table>
            </div>
        );
    }

}

export default Report