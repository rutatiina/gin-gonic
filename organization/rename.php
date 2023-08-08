<?php

$module_folder_name = "module_template";
$module_name = "wese/core/module_template";

$files = [
    "./controllers/controller_name.go",
    "./models/model_name.go",
    "./routes/route_name.go",
    "./go.mod",
];

$GoModuleName = "wese/core/organization";
$ControllerName = $ModelName = $RouteName = "Organization";
// $ModelName = "Organization";
// $RouteName = "Organization";
$RouteResourceName = "organizations";

foreach ($files as $file) {

    //read the entire string
    $str = file_get_contents($file);

    //replace controller name
    $str = str_replace("ControllerName", $ControllerName, $str);

    //replace model name
    $str = str_replace("ModelName", $ModelName, $str);

    //replace route name
    $str = str_replace("RouteName", $RouteName, $str);

    //replace route resource name
    $str = str_replace("resource-name", $RouteResourceName, $str);

    //replace route resource name
    $str = str_replace("wese/core/module_template", $GoModuleName, $str);

    //write the entire string
    file_put_contents($file, $str);
}

//rename files
if (file_exists("controllers/controller_name.go")) rename("controllers/controller_name.go", "controllers/organization.go");
if (file_exists("models/model_name.go")) rename("models/model_name.go", "models/organization.go");
if (file_exists("routes/route_name.go")) rename("routes/route_name.go", "routes/organization.go");

echo "Complete";
