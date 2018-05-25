import React, { Component } from 'react';
import GetReportXML from './GetRequest';
//import ReportRow  from './ReportRow';

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
                    <td>{row.Field1}</td>
                    <td>{row.Field2}</td>
                    <td>{row.Field3}</td>
                    <td>{row.Field4}</td>
                    <td>{row.Field5}</td>
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