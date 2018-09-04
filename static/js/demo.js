var Demo = (function() {
    function demoUpload() {
		var $uploadCrop;

		function readFile(input) {
 			if (input.files && input.files[0]) {
	            var reader = new FileReader();

	            reader.onload = function (e) {
					$('.upload-demo').addClass('ready');
	            	$uploadCrop.croppie('bind', {
	            		url: e.target.result
	            	}).then(function(){
	            		console.log('jQuery bind complete');
	            	});

	            }

	            reader.readAsDataURL(input.files[0]);
	        }
	        else {
		        swal("Sorry - you're browser doesn't support the FileReader API");
		    }
		}

		$uploadCrop = $('#upload-demo').croppie({
			viewport: {
				width: 100,
				height: 100,
				// type: 'circle'
			},
			enableExif: true
		});

		$('#upload').on('change', function () { readFile(this); });
		$('.upload-result').on('click', function (ev) {
			$uploadCrop.croppie('result', {
				type: 'canvas',
				size: 'viewport'
			}).then(function (resp) {
                console.log(resp);
                $.ajax({
                    url: "/upload-new",
                    type: "POST",
                    data: {imagebase64: resp},
                    success: function() {
                        alert("ok");
                    },
                    error: function() {
                         alert("error");
                    }
                });
			});
		});
	}

	function init() {
		demoUpload();
	}

	return {
		init: init
	};
}) ();
