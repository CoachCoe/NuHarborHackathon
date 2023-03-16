Hello and welcome to the coding part of the challenge.  This will be your guide for working with the data. Please reach out via the Discord channel if you have any questions or issues

There are two zipped files at this link:
https://drive.google.com/drive/folders/1vE_OHbpvNBNKMwv7938f3YZr7NWG4JlO?usp=share_link

All-CS.zip is all of the test/training data for CrowdStrike Falcon. The individual alerts are json objects and the file is a collection of those objects.  The file itself doesn't have a top level elelemt so itself it is not a valid jsob object

All-MS.zip is all of the test/training data for Microsoft Defender for Endpoint. The individual alerts are json objects and the file is a collection of those objects.  The file itself doesn't have a top level elelemt so itself it is not a valid jsob object

All alerts are key value pairs.  The exact format vaires from alert to alert but all alerts within each vendor do have a consistent set of core elecments 

key.xls shows the correct mapping of Microsoft Defender for Endpoiont alerts to CrowdStrike Falcon alerts

The MS.zip folder contains one version each of the indiviudal alerts for Microsoft Defender for Endpoint. Each alert has it's own json file. 

The CS.zip folder contains one version each of the indiviudal alerts for CrowdStrike Falcon. Each alert has it's own json file. 

Notes:
The alerts were generated and manipulated for the hackathon. You should not make any inferences about the distributiuon patterns. There is no logic to that and in the real-world the distribution is unpredicitable and will vary from organization to organization

Ignore all time stamps and event_id or AlertID's. There is no logic to them and no correlation to the ID's and the meaning of any alert. `

Becuase there was manual manipulation in creating the events, it's possible there are human errors, ie spelling or duplicate or missing info. These should be at a very minimum number

The intent is for this to be an NLP solution however it is possible some inferences and predictions can be made from the fomrat and/or structure of the alerts. It's also possible you will need multiple layers and/or a fine tuning layer on top of the core model 

In either case the goal is to create a model that learns from the alerts what the equivalent alert is in Microsoft Defender for Endpoint vs in CrowdStrike Falcon. The anticipated result is the common ontology, or common langauge to refer to and describe the same alert from different vendors 
