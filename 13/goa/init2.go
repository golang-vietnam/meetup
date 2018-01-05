	q := Q{
		comctrl.AddressCtrl{addressSvc, userSvc, *shipmentSvc}, // HL
		comctrl.BrandCtrl{conf, brandSvc},
		comctrl.CategoryCtrl{conf, categorySvc, *productSvc, tagSvc}, // HL
		ctrl.CompanyCtrl{companySvc},
		comctrl.FeedbackCtrl{feedbackSvc},
		ctrl.ImageCtrl{conf, imageSvc},
	}
