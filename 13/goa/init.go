addressSvc := svc.AddressSvc{store.AddressStore{cruder, lister}}
brandSvc := svc.BrandSvc{store.BrandStore{cruder, lister, beginer}}
categorySvc := svc.CategorySvc{store.CategoryStore{cruder, lister, beginer}}
companySvc := svc.CompanySvc{store.CompanyStore{cruder, lister, beginer}}
feedbackSvc := svc.FeedbackSvc{store.FeedbackStore{cruder, lister}}
imageSvc := svc.ImageSvc{store.ImageStore{conf, queryer}}
invoiceSvc := svc.InvoiceSvc{store.InvoiceStore{cruder, queryer, beginer, lister}}
contentSvc := svc.ContentSvc{store.ContentStore{cruder, lister, beginer}}
