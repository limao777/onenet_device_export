package structs

import (

)

type BackDevRet struct{
    Errno int	`json:"errno"`
    Data BackDevData
}

type BackDevData struct{
    Per_page int	`json:"per_page"`
    Page int	`json:"page"`
    Devices []BackDevDevices
}

type BackDevDevices struct{
    Id string	`json:"id"` 
    Private bool	`json:"private"`
    Title string	`json:"title"` 
    Desc string	`json:"desc"` 
    Tags interface{}	`json:"tags"` 
    Url string	`json:"url"` 
    Isdn string	`json:"isdn"` 
    Location interface{}	`json:"location"` 
    Protocol string	`json:"protocol"`
    Route_to interface{}	`json:"route_to"` 
    Auth_info interface{}	`json:"auth_info"` 
    Active_code string	`json:"active_code"` 
    Interval string	`json:"interval"` 
    Other interface{}	`json:"other"` 
    Key interface{}	`json:"key"` 
    Create_time string	`json:"create_time"` 

}